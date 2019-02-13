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
	"bufio"
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

type listCmd struct {
	fetchInfo bool
}

func (*listCmd) Name() string     { return "list" }
func (*listCmd) Synopsis() string { return "Parse RJ codes from stdin." }
func (*listCmd) Usage() string {
	return `Usage: list
Parse RJ codes from stdin.
`
}

func (c *listCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.fetchInfo, "fetch", false, "Also fetch info")
}

func (c *listCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := listMain(c); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func listMain(c *listCmd) error {
	if c.fetchInfo {
		c := dsutil.DefaultCache()
		defer c.Close()
		return mapCodes(os.Stdin, func(r dlsite.RJCode) error {
			w, err := dsutil.Fetch(c, r)
			if err != nil {
				return xerrors.Errorf("fetch work: %w", err)
			}
			printWork(os.Stdout, w)
			os.Stdout.Write([]byte("\n"))
			return nil
		})

	} else {
		return mapCodes(os.Stdin, func(r dlsite.RJCode) error {
			fmt.Println(r)
			return nil
		})
	}
}

func mapCodes(r io.Reader, f func(dlsite.RJCode) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		r := dlsite.Parse(s.Text())
		if err := f(r); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}
