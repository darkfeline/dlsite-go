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
	"log"
	"os"

	"github.com/google/subcommands"

	"go.felesatra.moe/dlsite/v2"
	"go.felesatra.moe/dlsite/v2/codes"
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
	df, err := dlsite.NewFetcher()
	if err != nil {
		log.Printf("Error making fetcher: %s", err)
		return subcommands.ExitFailure
	}
	defer df.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		c := codes.ParseRJCode(line)
		if c == "" {
			fmt.Println(line)
			continue
		}
		w, err := df.FetchWork(codes.WorkCode(c))
		if err != nil {
			log.Printf("Error fetching work info: %s", err)
			continue
		}
		fmt.Println(workInfoString(w))
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading stdin: %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func workInfoString(w *dlsite.Work) string {
	return fmt.Sprintf("%s [%s] %s", w.Code, w.Circle, w.Title)
}
