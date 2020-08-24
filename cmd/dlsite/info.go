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
	"log"
	"os"
	"strings"

	"github.com/google/subcommands"
	"go.felesatra.moe/dlsite/v2"
	"go.felesatra.moe/dlsite/v2/codes"
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
	r := codes.ParseRJCode(f.Arg(0))
	if r == "" {
		log.Printf("Invalid RJ code")
		return subcommands.ExitUsageError
	}
	if err := printInfo(r); err != nil {
		log.Printf("Error: %s\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func printInfo(c codes.RJCode) error {
	df, err := dlsite.NewFetcher()
	if err != nil {
		return fmt.Errorf("fetch work info: %w", err)
	}
	defer df.Close()
	w, err := df.FetchWork(codes.WorkCode(c))
	if err != nil {
		return fmt.Errorf("fetch work info: %w", err)
	}
	printWork(os.Stdout, w)
	return nil
}

func printWork(f io.Writer, w *dlsite.Work) (int, error) {
	const t = `%s
Title	%s
Circle	%s
Series	%s
CVs	%s
`
	return fmt.Fprintf(f, t, w.Code, w.Title, w.Circle, w.Series, strings.Join(w.CVs, ", "))
}
