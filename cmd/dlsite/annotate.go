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
	"os"

	"github.com/google/subcommands"

	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/dsutil"
)

type annotateCmd struct{}

func (*annotateCmd) Name() string     { return "annotate" }
func (*annotateCmd) Synopsis() string { return "Annotate DLSite RJ codes with work info." }
func (*annotateCmd) Usage() string {
	return `Usage: annotate
Annotate DLSite codes from stdin with work info.
`
}

func (*annotateCmd) SetFlags(f *flag.FlagSet) {}

func (c *annotateCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cache := dsutil.DefaultCache()
	defer cache.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		r := dlsite.Parse(line)
		if r == "" {
			fmt.Println(line)
			continue
		}
		w, err := dsutil.Fetch(cache, r)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error fetching work info:", err)
			continue
		}
		fmt.Println(workInfoString(w))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func workInfoString(w *dlsite.Work) string {
	return fmt.Sprintf("%s [%s] %s", w.RJCode, w.Maker, w.Name)
}
