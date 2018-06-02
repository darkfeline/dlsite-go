package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/subcommands"
)

func init() {
	commands = append(commands, subcommands.New("info", infoCmd))
}

func infoCmd(args []string) {
	if len(args) != 2 {
		infoUsage(os.Stderr)
		os.Exit(1)
	}
	c := dlsite.Parse(args[1])
	if c == "" {
		log.Fatal("Invalid RJ code")
	}
	w, err := dlsite.Fetch(c)
	if err != nil {
		log.Fatalf("Error fetching work info: %s", err)
	}
	const t = `%s
Name %s
Maker %s
Series %s
`
	fmt.Printf(t, w.RJCode, w.Name, w.Maker, w.Series)
}

func infoUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s info RJCODE\n", progName)
}
