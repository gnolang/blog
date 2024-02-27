package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gnolang/gno/tm2/pkg/commands"
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
		Exec: func(_ context.Context, args []string) error {
			return cfg.execPost(io, args)
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

func (cfg *cliCfg) execPost(io commands.IO, args []string) error {
	if len(args) > 1 {
		return errInvalidNumberOfArgs
	}

	// Stat passed in arg
	fileInfo, err := os.Stat(args[0])
	if err != nil {
		return errInvalidPath
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

	// Batch post request passed in with root argument
	if fileInfo.IsDir() {
		// Find file paths
		files, err := findFilePaths(args[0])
		if err != nil {
			return err
		}

		return cfg.post(client, files...)
	}

	// Single post request passed in an argument
	return cfg.post(client, args[0])
}

func (cfg *cliCfg) post(c gnoclient.Client, paths ...string) error {
	msgs := make([]gnoclient.MsgCall, 0, len(paths))

	// Go through post(s)
	for _, postPath := range paths {
		// Open file
		postFile, err := os.Open(postPath)
		if err != nil {
			return fmt.Errorf("cannot open file %q: %w", postPath, err)
		}

		// Parse post
		post, err := parsePost(postFile)
		if err != nil {
			return fmt.Errorf("cannot parse post %q: %w", postPath, err)
		}

		// Define function to call on the blog realm
		verb := "ModAddPost"

		// Check if post already exists on chain
		existsExpr := "PostExists(\"" + post.Slug + "\")"
		exists, _, err := c.QEval(cfg.BlogRealmPath, existsExpr)

		// If post exists, and user wants to edit it, use ModEditPost
		if err != nil && strings.Contains(exists, "true") && cfg.Edit {
			verb = "ModEditPost"
		}

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
