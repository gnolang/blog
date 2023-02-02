package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/peterbourgon/ff/ffcli"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		}
		os.Exit(1)
	}
}

type pushOpts struct {
	Debug   bool
	Publish bool
	Remote  string
	GnokeyOpts ...
}

func (opts *pushOpts) applyDefaults() {
	opts.Remote = "test3.gno.land:36657"
}

func (opts *pushOpts) flagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("blog push", flag.ExitOnError)
	fs.BoolVar(&opts.Debug, "d", false, "verbose output")
	fs.BoolVar(&opts.Publish, "p", false, "publish blogpost")
	return fs
}

func run(args []string) error {
	var opts rootOpts

	root := &ffcli.Command{
		ShortUsage: "blog <POST...>",
		FlagSet:    opts.flagSet(),
		Exec: func(ctx context.Context, args []string) error {
			return doRoot(ctx, args, opts)
		},
	}

	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}
	if err := root.Run(context.Background()); err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}
	return nil
}

type postMetadata struct {
	Title           string    `yaml:"title"`
	Tags            []string  `yaml:"tags"`
	Authors         []string  `yaml:"authors"`
	PublicationDate time.Time `yaml:"publication_date"`
	Slug            string    `yaml:"slug"`
}

func doRoot(ctx context.Context, posts []string, opts rootOpts) error {
	if len(args) < 1 {
		return flag.ErrHelp
	}

	for _, postPath := range posts {
		var matter postMetadata
		rest, err := frontmatter.MustParse(strings.NewReader(input), &matter)
		if err != nil {
			return fmt.Errorf("invalid post: %q: %w", postPath, err)
		}
		fmt.Printf("%+v\n", matter)
		fmt.Println(string(rest))
	}

	return nil
}
