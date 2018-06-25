/*
Command dlsite provides dlsite utilities.
*/
package main

import (
	"fmt"
	"io"
	"os"

	"go.felesatra.moe/subcommands"
)

var progName = os.Args[0]
var commands = make([]subcommands.Cmd, 0, 4)

func main() {
	if err := subcommands.Run(commands, os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, err)
		usage(os.Stderr)
		os.Exit(1)
	}
}

func usage(w io.Writer) {
	fmt.Fprintln(w, "Valid commands:")
	for _, c := range commands {
		fmt.Fprintln(w, c.Name())
	}
}
