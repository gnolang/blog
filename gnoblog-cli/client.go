package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gnolang/gno/pkgs/amino"
	"github.com/gnolang/gno/pkgs/crypto/keys"
	"github.com/gnolang/gno/pkgs/crypto/keys/client"
	"github.com/gnolang/gno/pkgs/sdk/vm"
	"github.com/gnolang/gno/pkgs/std"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		}
		os.Exit(1)
	}
}

type publishOpts struct {
	Debug   bool
	Publish bool
	PkgPath string

	// gnokey
	client.BaseOptions
	GasWanted       int64
	GasFee          string
	ChainID         string
	KeyNameOrBech32 string
}

func (opts *publishOpts) flagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("blog publish", flag.ExitOnError)
	fs.BoolVar(&opts.Debug, "debug", false, "verbose output")
	fs.BoolVar(&opts.Publish, "publish", false, "publish blogpost")
	fs.Int64Var(&opts.GasWanted, "gas-wanted", 1000, "gas requested for tx")
	fs.StringVar(&opts.GasFee, "gas-fee", "1ugnot", "gas payment fee")
	fs.StringVar(&opts.ChainID, "chainid", "test3", "")
	fs.StringVar(&opts.PkgPath, "pkgpath", "r/gnoland/blog", "blog realm path")

	// client.BaseOptions
	fs.StringVar(&opts.Home, "home", client.DefaultBaseOptions.Home, "home directory")
	fs.StringVar(&opts.Remote, "remote", "test3.gno.land:36657", "remote node URL")
	fs.BoolVar(&opts.Quiet, "quiet", false, "for parsing output")
	fs.BoolVar(&opts.InsecurePasswordStdin, "insecure-password-stdin", false, "WARNING! take password from stdin")
	return fs
}

func run(args []string) error {
	var opts publishOpts

	root := &ffcli.Command{
		ShortUsage: "gnoblog-cli KEY <POST...>",
		ShortHelp:  "(re)publish a blogpost",
		FlagSet:    opts.flagSet(),
		Exec: func(ctx context.Context, args []string) error {
			if len(args) < 2 {
				return flag.ErrHelp
			}
			opts.KeyNameOrBech32 = args[0]
			posts := args[1:]
			return doPublish(ctx, posts, opts)
		},
	}

	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	if !opts.Debug {
		log.SetOutput(ioutil.Discard)
	}

	if err := root.Run(context.Background()); err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}

	return nil
}

func doPublish(ctx context.Context, posts []string, opts publishOpts) error {
	log.Println("opts:", string(amino.MustMarshalJSON(opts)))

	kb, err := keys.NewKeyBaseFromDir(opts.Home)
	if err != nil {
		return err
	}
	info, err := kb.GetByNameOrAddress(opts.KeyNameOrBech32)
	if err != nil {
		return err
	}
	caller := info.GetAddress()
	log.Println("key:", caller)

	gasFee, err := std.ParseCoin(opts.GasFee)
	if err != nil {
		return fmt.Errorf("parse gas fee: %w", err)
	}

	for _, postPath := range posts {
		postFile, err := os.Open(postPath)
		if err != nil {
			return fmt.Errorf("open post %q: %w", postPath, err)
		}

		p, err := parsePost(postFile)
		if err != nil {
			return fmt.Errorf("parse post %q: %w", postPath, err)
		}

		// fmt.Printf("%+v\n", p)

		msg := vm.MsgCall{
			Caller:  caller,
			PkgPath: opts.PkgPath,
			Func:    "Render",
			Args:    []string{"p/" + p.Slug},
		}
		tx := std.Tx{
			Msgs:       []std.Msg{msg},
			Fee:        std.NewFee(opts.GasWanted, gasFee),
			Signatures: nil,
			// Memo:
		}

		log.Println(string(amino.MustMarshalJSON(tx)))
	}

	return nil
}
