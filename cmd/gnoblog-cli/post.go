package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	"github.com/gnolang/gno/tm2/pkg/commands"
)

type cliCfg struct {
	Publish       bool
	Edit          bool
	GasWanted     int64
	GasFee        string
	ChainId       string
	BlogRealmPath string

	KeyName               string
	GnoHome               string
	Remote                string
	Quiet                 bool
	InsecurePasswordStdIn bool
	ConfirmInput          bool
}

func newPostCommand(io commands.IO) *commands.Command {
	var (
		fs  = flag.NewFlagSet("Post", flag.ExitOnError)
		cfg = &cliCfg{}
	)

	// Register flagset
	cfg.RegisterFlags(fs)

	// Make Post command
	return commands.NewCommand(
		commands.Metadata{
			Name:       "post",
			ShortUsage: "post <FILE OR FILES_DIR> [flags]",
			LongHelp:   `Post one or more files. Passing in a file will post that single file, while passing in a directory will search for all README.md files and batch post them.`,
		},
		cfg,
		func(_ context.Context, args []string) error {
			return execPost(io, args, cfg)
		},
	)
}

// RegisterFlags registers flags for cliCfg
func (cfg *cliCfg) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.KeyName,
		"key",
		"",
		"name of keypair to use for deployment",
	)
	fs.BoolVar(&cfg.Publish,
		"publish",
		false,
		"publish blogpost",
	)
	fs.BoolVar(&cfg.Edit,
		"edit",
		false,
		"edit mode",
	)
	fs.Int64Var(&cfg.GasWanted,
		"gas-wanted",
		5000000,
		"gas requested for tx",
	)
	fs.StringVar(&cfg.GasFee,
		"gas-fee",
		"1000000ugnot",
		"gas payment fee",
	)
	fs.StringVar(&cfg.ChainId,
		"chainid",
		"dev",
		"chain ID",
	)
	fs.StringVar(&cfg.BlogRealmPath,
		"pkgpath",
		"gno.land/r/gnoland/blog",
		"blog realm path",
	)
	fs.StringVar(&cfg.GnoHome,
		"home",
		gnoenv.HomeDir(),
		"home directory",
	)
	fs.StringVar(&cfg.Remote,
		"remote",
		"localhost:26657",
		"remote node URL",
	)
	fs.BoolVar(&cfg.InsecurePasswordStdIn,
		"insecure-password-stdin",
		false,
		"WARNING! take password from stdin",
	)
	fs.BoolVar(&cfg.ConfirmInput,
		"confirm-input",
		false,
		"ask user to confirm input",
	)
}

func execPost(io commands.IO, args []string, cfg *cliCfg) error {
	if len(args) != 1 {
		return ErrInvalidNumberOfArgs
	}

	if cfg.KeyName == "" {
		return ErrEmptyKeyName
	}

	// Ask user for confirming the ChainID
	if cfg.ConfirmInput && !askForConfirmation(cfg.ChainId, cfg.KeyName) {
		fmt.Println("operation canceled by the user, exiting")
		return nil
	}

	// Stat passed in arg
	fileInfo, err := os.Stat(args[0])
	if err != nil {
		return fmt.Errorf("unable to stat %q: %w", args[0], err)
	}

	var pass string
	if cfg.Quiet {
		pass, err = io.GetPassword("", cfg.InsecurePasswordStdIn)
	} else {
		pass, err = io.GetPassword("Enter password:", cfg.InsecurePasswordStdIn)
	}
	if err != nil {
		return err
	}

	// Initialize signer
	signer, err := initSigner(cfg, pass)
	if err != nil {
		return err
	}

	// Initialize Gnoclient
	rpc, err := initRPCClient(cfg)
	if err != nil {
		return err
	}

	client := gnoclient.Client{
		Signer:    signer,
		RPCClient: rpc,
	}

	// Batch Post request passed in with root argument
	if fileInfo.IsDir() {
		// Find file paths
		files, err := findFilePaths(args[0])
		if err != nil {
			return err
		}

		return post(client, cfg, files...)
	}

	// Single Post request passed in an argument
	return post(client, cfg, args[0])
}

func post(c gnoclient.Client, cfg *cliCfg, paths ...string) error {
	msgs := make([]vm.MsgCall, 0, len(paths))

	// Get account info
	account, err := c.Signer.Info()
	if err != nil {
		return fmt.Errorf("getting signer info failed: %w", err)
	}

	address := account.GetAddress()

	signingAcc, _, err := c.QueryAccount(address)
	if err != nil {
		return fmt.Errorf("query account %q failed: %w", account, err)
	}

	nonce := signingAcc.GetSequence()
	accNumber := signingAcc.GetAccountNumber()

	// Set up base config
	baseTxCfg := gnoclient.BaseTxCfg{
		GasFee:         cfg.GasFee,
		GasWanted:      cfg.GasWanted,
		AccountNumber:  accNumber,
		SequenceNumber: nonce,
		Memo:           "Posted from gnoblog-cli",
	}

	// Save title for printing to cli
	var postTitle string
	// Go through Post(s)
	for _, postPath := range paths {
		// Open file
		postFile, err := os.Open(postPath)
		if err != nil {
			return fmt.Errorf("cannot open file %q: %w", postPath, err)
		}

		// Parse Post
		post, err := parsePost(postFile)
		if err != nil {
			fmt.Printf("skipping post at %q, cannot parse: %v\n", postPath, err)
			continue
		}

		// Define function to call on the blog realm
		verb := "ModAddPost"

		// Check if Post already exists on chain
		existsExpr := "PostExists(\"" + post.Slug + "\")"
		exists, _, err := c.QEval(cfg.BlogRealmPath, existsExpr)
		if err != nil {
			slog.Error("error while checking if Post exists", "error", err, "slug", post.Slug)
		}

		bExists := strings.Contains(exists, "true")
		if cfg.Edit {
			if !bExists {
				return fmt.Errorf("%s is not on chain yet - disable the edit flag\n", post.Title)
			}
			// If Post exists, and user wants to edit it, use ModEditPost
			verb = "ModEditPost"
		} else if bExists {
			// if a post is already on chain, and we are not editing it, just skip it
			// otherwise, batch transactions will fail if a single MsgCall fails
			continue
		}

		callMsg := vm.MsgCall{
			Caller:  address,
			Send:    nil,
			PkgPath: cfg.BlogRealmPath,
			Func:    verb,
			Args: []string{
				post.Slug,
				post.Title,
				post.Body,
				post.PublicationDate.Format(time.RFC3339),
				strings.Join(post.Authors, ","),
				strings.Join(post.Tags, ","),
			},
		}

		msgs = append(msgs, callMsg)
		postTitle = post.Title
	}

	if len(msgs) == 0 {
		return fmt.Errorf("%w, exiting", ErrNoNewOrChangedPosts)
	}

	_, err = c.Call(baseTxCfg, msgs...)
	if err != nil {
		return err
	}

	// Print success messages
	action := "posted"
	if cfg.Edit {
		action = "edited"
	}

	if len(msgs) == 1 {
		fmt.Printf("Successfully %s \"%s\"!\n", action, postTitle)
	} else {
		fmt.Printf("Successfully %s %d posts!\n", action, len(msgs))
	}

	return nil
}

func parsePost(reader io.Reader) (*Post, error) {
	var p Post
	rest, err := frontmatter.MustParse(reader, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid Post frontmatter: %w", err)
	}

	body := string(rest)
	p.Title, err = extractTitle(body)
	if err != nil {
		return nil, err
	}

	p.Body = removeTitle(body, p.Title)

	if len(p.Tags) != 0 {
		p.Tags = removeWhitespace(p.Tags)
	}

	if p.PublicationDate == nil {
		now := time.Now()
		p.PublicationDate = &now
	}

	return &p, nil
}
