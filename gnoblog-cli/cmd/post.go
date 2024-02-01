package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
	"strings"
	"time"
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

func newPostCommand() *ffcli.Command {
	var (
		fs  = flag.NewFlagSet("post", flag.ExitOnError)
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
		Exec:       cfg.execPost,
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

func (cfg *cliCfg) execPost(_ context.Context, args []string) error {
	if len(args) > 1 {
		fmt.Println(args)

		return errInvalidNumberOfArgs
	}

	fileInfo, err := os.Stat(args[0])
	if err != nil {
		return errInvalidPath
	}

	// Init IO for password input
	io := commands.NewDefaultIO()

	var pass string
	if cfg.Quiet {
		pass, err = io.GetPassword("", cfg.InsecurePasswordStdIn)
	} else {
		pass, err = io.GetPassword("Enter password:", cfg.InsecurePasswordStdIn)
	}
	if err != nil {
		return err
	}

	// Validate chainID that was passed in
	//if err = validateChainID(cfg.ChainId); err != nil {
	//	return err
	//}

	// Initialize Gnoclient
	client := gnoclient.Client{
		Signer:    initSigner(cfg, pass),
		RPCClient: initRPCClient(cfg),
	}

	// Batch post request passed in
	if fileInfo.IsDir() {
		return cfg.batchPost(client, args[0])
	}

	// Single post request passed in an argument
	return cfg.singlePost(client, args[0])
}

func (cfg *cliCfg) singlePost(c gnoclient.Client, postPath string) error {
	postFile, err := os.Open(postPath)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %w", postPath, err)
	}

	post, err := parsePost(postFile)
	if err != nil {
		return fmt.Errorf("cannot parse post %q: %w", postPath, err)
	}

	signingAcc, _, err := c.QueryAccount(c.Signer.Info().GetAddress())
	if err != nil {
		return err
	}

	nonce := signingAcc.GetSequence()
	accNumber := signingAcc.GetAccountNumber()

	// Check if post already exists
	existsExpr := "PostExists(\"" + post.Slug + "\")"
	exists, _, err := c.QEval("gno.land/r/gnoland/blog", existsExpr)

	verb := "ModAddPost"
	if strings.Contains(exists, "true") && cfg.Edit {
		verb = "ModEditPost"
	}

	msg := gnoclient.MsgCall{
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

	baseTxCfg := gnoclient.BaseTxCfg{
		GasFee:         cfg.GasFee,
		GasWanted:      cfg.GasWanted,
		AccountNumber:  accNumber,
		SequenceNumber: nonce,
		Memo:           "Posted from gnoblog-cli",
	}

	if !cfg.Publish {
		fmt.Println("--publish flag not set to true, exiting")
		// todo see how to display generated data
		return nil
	}

	fmt.Println(msg)
	_, err = c.Call(baseTxCfg, msg)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully posted %s", post.Title)

	return nil
}

func (cfg *cliCfg) batchPost(c gnoclient.Client, dirPath string) error {
	files, err := findFilePaths(dirPath)
	if err != nil {
		return err
	}

	msgs := make([]gnoclient.MsgCall, 0, len(files))
	for _, postPath := range files {
		postFile, err := os.Open(postPath)
		if err != nil {
			return fmt.Errorf("cannot open file %q: %w", postPath, err)
		}

		post, err := parsePost(postFile)
		if err != nil {
			return fmt.Errorf("cannot parse post %q: %w", postPath, err)
		}

		callCfg := gnoclient.MsgCall{
			PkgPath:  cfg.BlogRealmPath,
			FuncName: "ModAddPost",
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

	signingAcc, _, err := c.QueryAccount(c.Signer.Info().GetAddress())
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
