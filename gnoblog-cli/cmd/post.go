package main

import (
	"context"
	"flag"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/peterbourgon/ff/v3/ffcli"
)

type postCfg struct {
	base     *baseCliCfg
	postPath string
}

func newPostCommand(base *baseCliCfg) *ffcli.Command {
	cfg := &postCfg{
		base: base,
	}
	fs := flag.NewFlagSet("batch", flag.ExitOnError)

	cfg.registerFlags(fs)

	return &ffcli.Command{
		Name:       "post",
		ShortUsage: "post [flags]",
		LongHelp:   "post a single post",
		FlagSet:    fs,
		Exec:       cfg.exec,
	}
}

func (p *postCfg) registerFlags(fs *flag.FlagSet) {
	fs.StringVar(&p.postPath,
		"path",
		"",
		"path to the post README",
	)
}

func (p *postCfg) exec(_ context.Context, _ []string) error {
	if p.postPath == "" {
		return errors.New("post path cannot be empty")
	}

	return nil
}
