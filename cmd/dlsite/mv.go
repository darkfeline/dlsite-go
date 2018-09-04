package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/subcommands"
	"github.com/pkg/errors"

	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/dsutil"
)

type mvCmd struct {
}

func (*mvCmd) Name() string     { return "mv" }
func (*mvCmd) Synopsis() string { return "Rename work dirs using DLSite info." }
func (*mvCmd) Usage() string {
	return `Usage: mv dir [rjcode]
Rename work dirs using DLSite info.
`
}

func (*mvCmd) SetFlags(f *flag.FlagSet) {

}

func (c *mvCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() < 1 {
		fmt.Fprint(os.Stderr, c.Usage())
		return subcommands.ExitUsageError
	}
	path := f.Arg(0)
	var s string
	if f.NArg() > 1 {
		s = f.Arg(1)
	} else {
		s = path
	}
	r := dlsite.Parse(s)
	if r == "" {
		fmt.Fprintf(os.Stderr, "invalid RJ code %s\n", s)
		return subcommands.ExitUsageError
	}
	if err := mvMain(path, r); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func mvMain(path string, r dlsite.RJCode) error {
	c := dsutil.DefaultCache()
	defer c.Close()
	w, err := dsutil.Fetch(c, r)
	if err != nil {
		return errors.Wrap(err, "fetch work info")
	}
	new := filepath.Join(filepath.Dir(path), workFilename(w))
	if new == path {
		return nil
	}
	if err := os.Rename(path, new); err != nil {
		return err
	}
	return nil
}

func workFilename(w *dlsite.Work) string {
	p := fmt.Sprintf("%s [%s] %s", w.RJCode, w.Maker, w.Name)
	return escapeFilename(p)
}

func escapeFilename(p string) string {
	return strings.Replace(p, "/", "_", -1)
}
