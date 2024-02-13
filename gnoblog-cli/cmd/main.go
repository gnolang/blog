package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	// Create the root command
	cmd := &ffcli.Command{
		ShortUsage: "<subcommand> [flags] [<arg>...]",
		LongHelp:   "The CLI for easy use of the r/blog realm",
		FlagSet:    nil,
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	// Add the subcommands
	cmd.Subcommands = []*ffcli.Command{
		newPostCommand(),
	}

	// Run root command
	if err := cmd.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
