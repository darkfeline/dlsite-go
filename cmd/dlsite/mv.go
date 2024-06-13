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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/subcommands"

	"go.felesatra.moe/dlsite/v2"
	"go.felesatra.moe/dlsite/v2/codes"
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
	r := codes.ParseCode(s)
	if r == "" {
		log.Printf("Invalid RJ code %s", s)
		return subcommands.ExitUsageError
	}
	if err := mvMain(path, r); err != nil {
		log.Printf("Error: %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func mvMain(path string, r codes.WorkCode) error {
	df, err := dlsite.NewFetcher()
	if err != nil {
		return fmt.Errorf("fetch work info: %w", err)
	}
	defer df.Close()
	w, err := df.FetchWork(r)
	if err != nil {
		return fmt.Errorf("fetch work info: %w", err)
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
	p := workInfoString(w)
	return escapeFilename(p)
}

func escapeFilename(p string) string {
	return strings.Replace(p, "/", "_", -1)
}
