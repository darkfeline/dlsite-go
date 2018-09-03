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
	p := f.Arg(0)
	var r dlsite.RJCode
	if f.NArg() > 1 {
		r = dlsite.Parse(f.Arg(1))
	} else {
		r = dlsite.Parse(p)
	}
	if err := mvMain(p, r); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func mvMain(p string, r dlsite.RJCode) error {
	c := dsutil.DefaultCache()
	defer c.Close()
	w, err := dsutil.Fetch(c, r)
	if err != nil {
		return errors.Wrap(err, "fetch work info")
	}
	new := filepath.Join(filepath.Dir(p), workFilename(w))
	if err := os.Rename(p, new); err != nil {
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
