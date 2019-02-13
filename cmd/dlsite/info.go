// Copyright (C) 2018  Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/subcommands"
	"golang.org/x/xerrors"

	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/dsutil"
)

type infoCmd struct {
}

func (*infoCmd) Name() string     { return "info" }
func (*infoCmd) Synopsis() string { return "Show info for work." }
func (*infoCmd) Usage() string {
	return `Usage: info rjcode
Show info for work.
`
}

func (*infoCmd) SetFlags(f *flag.FlagSet) {

}

func (c *infoCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() != 1 {
		fmt.Fprint(os.Stderr, c.Usage())
		return subcommands.ExitUsageError
	}
	r := dlsite.Parse(f.Arg(0))
	if r == "" {
		fmt.Fprintf(os.Stderr, "Invalid RJ code\n")
		return subcommands.ExitUsageError
	}
	if err := printInfo(r); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func printInfo(r dlsite.RJCode) error {
	c := dsutil.DefaultCache()
	defer c.Close()
	w, err := dsutil.Fetch(c, r)
	if err != nil {
		return xerrors.Errorf("fetch work info: %w", err)
	}
	printWork(os.Stdout, w)
	return nil
}

func printWork(f io.Writer, w *dlsite.Work) (int, error) {
	const t = `%s
Name %s
Maker %s
Series %s
`
	return fmt.Fprintf(f, t, w.RJCode, w.Name, w.Maker, w.Series)
}
