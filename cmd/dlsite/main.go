// Command dlsite provides dlsite utilities.
package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&infoCmd{}, "")
	subcommands.Register(&listCmd{}, "")
	subcommands.Register(&mvCmd{}, "")
	subcommands.Register(&orgCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
