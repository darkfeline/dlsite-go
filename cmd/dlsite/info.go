package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/dsutil"
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
	r := dlsite.Parse(args[1])
	if r == "" {
		log.Fatal("Invalid RJ code")
	}
	if err := printInfo(r); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func printInfo(r dlsite.RJCode) error {
	c := dsutil.DefaultCache()
	defer c.Close()
	w, err := dsutil.Fetch(c, r)
	if err != nil {
		return errors.Wrap(err, "fetch work info")
	}
	const t = `%s
Name %s
Maker %s
Series %s
`
	fmt.Printf(t, w.RJCode, w.Name, w.Maker, w.Series)
	return nil
}

func infoUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s info RJCODE\n", progName)
}
