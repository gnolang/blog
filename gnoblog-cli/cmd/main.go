package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
)

type baseCliCfg struct {
	debug     bool
	publish   bool
	edit      bool
	gasWanted int64
	gasFee    string
	chainId   string
	blogPath  string

	keyName               string
	gnoHome               string
	remote                string
	quiet                 bool
	insecurePasswordStdIn bool
}

func main() {
	var (
		baseCfg = &baseCliCfg{}
		fs      = flag.NewFlagSet("root", flag.ExitOnError)
	)

	// Register the flags
	baseCfg.registerFlags(fs)

	// Create the root command
	cmd := &ffcli.Command{
		ShortUsage: "<subcommand> [flags] [<arg>...] <KEY>",
		LongHelp:   "The CLI for easy use of the r/blog realm",
		FlagSet:    fs,
		Exec: func(_ context.Context, args []string) error {
			if len(args) != 1 {
				return errors.New("deployment key not specified")
			}
			baseCfg.keyName = args[0]
			return flag.ErrHelp
		},
	}

	// Add the subcommands
	cmd.Subcommands = []*ffcli.Command{
		newPostCommand(baseCfg),
		newBatchCommand(baseCfg),
	}

	if err := cmd.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v", err)

		os.Exit(1)
	}
}

func (cfg *baseCliCfg) registerFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.publish, "publish", false, "publish blogpost")
	fs.BoolVar(&cfg.edit, "edit", false, "edit mode")
	fs.Int64Var(&cfg.gasWanted, "gas-wanted", 2000000, "gas requested for tx")
	fs.StringVar(&cfg.gasFee, "gas-fee", "1000000ugnot", "gas payment fee")
	fs.StringVar(&cfg.chainId, "chainid", "dev", "chain ID")
	fs.StringVar(&cfg.blogPath, "pkgpath", "gno.land/r/gnoland/blog", "blog realm path")

	// keysclient.BaseOptions
	fs.StringVar(&cfg.gnoHome, "home", "/Users/sasurai/gnohome", "home directory")
	fs.StringVar(&cfg.remote, "remote", "localhost:26657", "remote node URL")
	fs.BoolVar(&cfg.quiet, "quiet", false, "for parsing output")
	fs.BoolVar(&cfg.insecurePasswordStdIn, "insecure-password-stdin", false, "WARNING! take password from stdin")
}
