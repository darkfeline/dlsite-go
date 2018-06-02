package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/dsutil"
	"go.felesatra.moe/subcommands"
)

func init() {
	commands = append(commands, subcommands.New("mv", mvCmd))
}

func mvCmd(args []string) {
	if len(args) < 2 {
		mvUsage(os.Stderr)
		os.Exit(1)
	}
	n := args[1]
	var r dlsite.RJCode
	if len(args) > 3 {
		r = dlsite.Parse(args[2])
	} else {
		r = dlsite.Parse(n)
	}
	if err := mvMain(n, r); err != nil {
		log.Fatal(err)
	}
}

func mvMain(n string, r dlsite.RJCode) error {
	c := dsutil.DefaultCache()
	defer c.Close()
	w, err := dsutil.Fetch(c, r)
	if err != nil {
		return errors.Wrap(err, "fetch work info")
	}
	if err := os.Rename(n, filepath.Join(filepath.Dir(n), workFilename(w))); err != nil {
		return err
	}
	return nil
}

func workFilename(w *dlsite.Work) string {
	p := fmt.Sprintf("%s [%s] %s", w.RJCode, w.Maker, w.Name)
	return escapeFilename(p)
}

func escapeFilename(p string) string {
	return strings.Replace("/", "_", p, -1)
}

func mvUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s mv FILE [RJCODE]\n", progName)
}
