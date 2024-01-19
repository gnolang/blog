package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
)

type postCfg struct {
	*baseCliCfg
	path string
}

func newPostCommand(base *baseCliCfg) *ffcli.Command {
	cfg := &postCfg{
		baseCliCfg: base,
	}
	fs := flag.NewFlagSet("post", flag.ExitOnError)

	cfg.registerFlags(fs)

	return &ffcli.Command{
		Name:       "post",
		ShortUsage: "post [flags]",
		LongHelp:   "post a single post",
		FlagSet:    fs,
		Exec:       cfg.execPost,
	}
}

func (cfg *postCfg) registerFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.path,
		"path",
		"",
		"path to post README. if a dir is passed in, a batch post happens",
	)
}

func (cfg *postCfg) execPost(_ context.Context, _ []string) error {
	if cfg.path == "" {
		return errors.New("post path cannot be empty")
	}

	// Init IO for password input
	io := commands.NewDefaultIO()

	var pass string
	var err error
	if cfg.quiet {
		pass, err = io.GetPassword("", cfg.insecurePasswordStdIn)
	} else {
		pass, err = io.GetPassword("Enter password:", cfg.insecurePasswordStdIn)
	}

	if err != nil {
		return err
	}

	// Validate chainID that was passed in
	if err := validateChainID(cfg.chainId); err != nil {
		return err
	}

	client := gnoclient.Client{
		Signer:    initSigner(cfg, pass),
		RPCClient: initRPCClient(cfg),
	}

	//// Batch post request passed in
	//if cfg.batchRoot != "" {
	//	return cfg.batchPost()
	//}

	// Single post request passed in
	return cfg.singlePost(client)
}

func (cfg *postCfg) singlePost(c gnoclient.Client) error {
	postFile, err := os.Open(cfg.path)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %w", cfg.path, err)
	}

	post, err := parsePost(postFile)
	if err != nil {
		return fmt.Errorf("cannot parse post %q: %w", cfg.path, err)
	}

	signingAcc, _, err := c.QueryAccount(c.Signer.Info().GetAddress())

	if err != nil {
		return err
	}

	nonce := signingAcc.GetSequence()
	accNumber := signingAcc.GetAccountNumber()

	// add check if post exists to call edit

	callCfg := gnoclient.CallCfg{
		PkgPath:  cfg.blogPath,
		FuncName: "ModAddPost",
		Args: []string{
			post.Slug,
			post.Title,
			post.Body,
			"tags",
		},
		GasFee:         cfg.gasFee,
		GasWanted:      cfg.gasWanted,
		Send:           "",
		AccountNumber:  accNumber,
		SequenceNumber: nonce,
		Memo:           "Posted from gnoblog-cli",
	}

	call, err := c.Call(callCfg)
	if err != nil {
		return err
	}

	fmt.Printf("Success! %s", call.DeliverTx.Info)

	return nil
}

func (cfg *postCfg) batchPost() error {
	return nil
}
