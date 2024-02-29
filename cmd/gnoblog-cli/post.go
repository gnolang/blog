package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/adrg/frontmatter"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	"github.com/peterbourgon/ff/v3/ffcli"
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
}

func newPostCommand(io commands.IO) *ffcli.Command {
	var (
		fs  = flag.NewFlagSet("Post", flag.ExitOnError)
		cfg = &cliCfg{}
	)

	// Register flagset
	cfg.registerFlags(fs)

	// Make Post command
	return &ffcli.Command{
		Name:       "post",
		ShortUsage: "post <FILE OR FILES_DIR> [flags]",
		LongHelp:   `Post one or more files. Passing in a file will post that single file, while passing in a directory will search for all .md files and batch post them.`,
		FlagSet:    fs,
		Exec: func(_ context.Context, args []string) error {
			return execPost(io, args, cfg)
		},
	}
}

// Register flags
func (cfg *cliCfg) registerFlags(fs *flag.FlagSet) {
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
}

func execPost(io commands.IO, args []string, cfg *cliCfg) error {
	if len(args) > 1 {
		return ErrInvalidNumberOfArgs
	}

	if cfg.KeyName == "" {
		return ErrEmptyKeyName
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
	client := gnoclient.Client{
		Signer:    signer,
		RPCClient: initRPCClient(cfg),
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
	msgs := make([]gnoclient.MsgCall, 0, len(paths))

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
			return fmt.Errorf("cannot parse Post %q: %w", postPath, err)
		}

		// Define function to call on the blog realm
		verb := "ModAddPost"

		// Check if Post already exists on chain
		existsExpr := "PostExists(\"" + post.Slug + "\")"
		exists, _, err := c.QEval(cfg.BlogRealmPath, existsExpr)
		if err != nil {
			slog.Error("error while checking if Post exists", "error", err, "slug", post.Slug)
		}

		// If Post exists, and user wants to edit it, use ModEditPost
		if cfg.Edit && strings.Contains(exists, "true") {
			verb = "ModEditPost"
		}

		// Pack calls to the chain
		callCfg := gnoclient.MsgCall{
			PkgPath:  cfg.BlogRealmPath,
			FuncName: verb,
			Args: []string{
				post.Slug,
				post.Title,
				post.Body,
				post.PublicationDate.Format(time.RFC3339),
				strings.Join(post.Authors, ","),
				strings.Join(post.Tags, ","),
			},
		}
		msgs = append(msgs, callCfg)
	}

	account := c.Signer.Info().GetAddress()
	signingAcc, _, err := c.QueryAccount(account)
	if err != nil {
		return fmt.Errorf("query account %q failed: %w", account, err)
	}

	nonce := signingAcc.GetSequence()
	accNumber := signingAcc.GetAccountNumber()

	baseTxCfg := gnoclient.BaseTxCfg{
		GasFee:         cfg.GasFee,
		GasWanted:      cfg.GasWanted,
		AccountNumber:  accNumber,
		SequenceNumber: nonce,
		Memo:           "Posted from gnoblog-cli",
	}

	_, err = c.Call(baseTxCfg, msgs...)
	if err != nil {
		return err
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
