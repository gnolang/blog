package main

import (
	"context"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"os"
)

func main() {
	io := commands.NewDefaultIO()
	cmd := newRootCmd(io)

	cmd.Execute(context.Background(), os.Args[1:])
}

func newRootCmd(io commands.IO) *commands.Command {
	cmd := commands.NewCommand(
		commands.Metadata{
			ShortUsage: "<subcommand> [flags] [<arg>...]",
			LongHelp:   "The CLI for easy use of the r/blog realm",
		},
		commands.NewEmptyConfig(),
		commands.HelpExec,
	)

	cmd.AddSubCommands(
		newPostCommand(io),
	)

	return cmd
}
