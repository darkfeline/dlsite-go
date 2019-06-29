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
	"math/rand"
	"os"
	"time"

	"github.com/google/subcommands"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/cache"
	"golang.org/x/xerrors"
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
	ch, err := cache.OpenDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return subcommands.ExitFailure
	}
	defer ch.Close()
	ks, err := ch.Keys()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return subcommands.ExitFailure
	}
	rand.Seed(time.Now().UnixNano())
	const sleepMax = 2 * time.Second
	for _, k := range ks {
		time.Sleep(time.Duration(rand.Int63n(int64(sleepMax))))
		fmt.Fprintf(os.Stderr, "Refreshing %s\n", k)
		if err := refreshWork(ch, k); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}
	}
	return subcommands.ExitSuccess
}

func refreshWork(c *cache.Cache, r dlsite.RJCode) error {
	w, err := dlsite.Fetch(r)
	if err != nil {
		return xerrors.Errorf("refresh work %v: %w", r, err)
	}
	if err := c.Put(w); err != nil {

		return xerrors.Errorf("refresh work %v: %w", r, err)

	}
	return nil
}
