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
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/google/subcommands"
	"go.felesatra.moe/dlsite/v2"
	"go.felesatra.moe/dlsite/v2/codes"
)

type refreshCmd struct {
}

func (*refreshCmd) Name() string     { return "refresh" }
func (*refreshCmd) Synopsis() string { return "Refresh cache." }
func (*refreshCmd) Usage() string {
	return `Usage: refresh
Refresh the DLSite info cache.
`
}

func (*refreshCmd) SetFlags(f *flag.FlagSet) {}

func (c *refreshCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := refreshWorks(os.Stdin); err != nil {
		log.Printf("Error: %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func refreshWorks(r io.Reader) error {
	df, err := dlsite.NewFetcher()
	if err != nil {
		return err
	}
	defer df.Close()
	rand.Seed(time.Now().UnixNano())
	const sleepMax = 2 * time.Second
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		c := codes.ParseCode(line)
		if c == "" {
			log.Printf("Bad line %s", line)
			continue
		}
		time.Sleep(time.Duration(rand.Int63n(int64(sleepMax))))
		log.Printf("Refreshing %s", c)
		if _, err := df.FetchWorkDirectly(c); err != nil {
			log.Printf("Error: %s", err)
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
