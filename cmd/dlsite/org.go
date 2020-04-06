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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
	"go.felesatra.moe/dlsite/v2"
	"go.felesatra.moe/dlsite/v2/codes"
)

type orgCmd struct {
	dry  bool
	all  bool
	desc bool
}

func (*orgCmd) Name() string     { return "org" }
func (*orgCmd) Synopsis() string { return "Organize works." }
func (*orgCmd) Usage() string {
	return `Usage: org [dir]
Organize works.
`
}

func (c *orgCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.dry, "dryrun", false, "Dry run")
	f.BoolVar(&c.all, "all", false, "Organize all works recursively")
	f.BoolVar(&c.desc, "descriptions", false, "Fetch work descriptions")
}

func (c *orgCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var dir string
	switch f.NArg() {
	case 0:
		var err error
		dir, err = os.Getwd()
		if err != nil {
			log.Printf("Could not get current directory: %s", err)
			return subcommands.ExitFailure
		}
	case 1:
		dir = f.Arg(0)
	default:
		fmt.Fprint(os.Stderr, c.Usage())
		return subcommands.ExitUsageError
	}
	if err := orgMain(dir, c.dry, c.all, c.desc); err != nil {
		log.Printf("Error: %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func orgMain(dir string, dry, all, desc bool) error {
	var w []relPath
	var err error
	if all {
		w, err = findAllWorks(dir)
	} else {
		w, err = findWorks(dir)
	}
	if err != nil {
		return fmt.Errorf("find works: %w", err)
	}
	df, err := dlsite.NewFetcher()
	if err != nil {
		return fmt.Errorf("fetch work info: %w", err)
	}
	defer df.Close()
	for _, w := range w {
		log.Printf("Organizing %s", w)
		err := organizeWork(df, dir, w, dry, desc)
		if err != nil {
			return err
		}
	}
	return nil
}

// relPath is the path of a work relative to the organize root directory.
type relPath string

func (p relPath) relativeTo(root string) string {
	return filepath.Join(root, string(p))
}

// findWorks returns the relative paths for works found in the directory.
func findWorks(dir string) ([]relPath, error) {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var w []relPath
	for _, fi := range fi {
		if codes.ParseRJCode(fi.Name()) != "" {
			w = append(w, relPath(fi.Name()))
		}
	}
	return w, nil
}

// findAllWorks returns the relative paths for works found in the directory recursively.
func findAllWorks(dir string) ([]relPath, error) {
	var w []relPath
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking %s: %s", path, err)
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		if codes.ParseRJCode(info.Name()) != "" {
			// path has dir as a prefix plus a slash: dir/some/path
			p := path[len(dir)+1:]
			w = append(w, relPath(p))
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk: %w", err)
	}
	return w, nil
}

// organizeWork organizes the work with the given path relative to the
// top directory.  This involves moving it to the correct path and
// possibly adding some files.
func organizeWork(f *dlsite.Fetcher, topdir string, p relPath, dry, desc bool) error {
	w, err := getDirWork(f, string(p))
	if err != nil {
		return fmt.Errorf("get work info: %w", err)
	}
	new := workPath(w)
	if dry {
		if new != p {
			log.Printf("Would rename %s to %s", p, new)
		}
		return nil
	}
	if desc {
		log.Printf("Adding description files for %s", p)
		if err := addDLSiteFiles(w, p.relativeTo(topdir)); err != nil {
			return fmt.Errorf("add desc files: %w", err)
		}
	}
	if new == p {
		return nil
	}
	log.Printf("Moving %s", p)
	if err := renameWork(topdir, p, new); err != nil {
		return fmt.Errorf("rename work: %w", err)
	}
	return nil
}

// workPath returns the desired path for a work.
func workPath(w *dlsite.Work) relPath {
	// Empty parts are ignored by Join.
	return relPath(filepath.Join(escapeFilename(w.Circle), escapeFilename(w.Series),
		escapeFilename(fmt.Sprintf("%s %s", w.Code, w.Title))))
}

func renameWork(top string, old, new relPath) error {
	oldp := old.relativeTo(top)
	newp := new.relativeTo(top)
	if err := os.MkdirAll(filepath.Dir(newp), 0777); err != nil {
		return err
	}
	if err := os.Rename(oldp, newp); err != nil {
		return err
	}
	return nil
}

const (
	descFile = "dlsite-description.txt"
)

func addDLSiteFiles(w *dlsite.Work, p string) error {
	if w.Description != "" {
		fp := filepath.Join(p, descFile)
		ok, err := fileExists(fp)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		if err := ioutil.WriteFile(fp, []byte(w.Description), 0666); err != nil {
			return err
		}
	}
	return nil
}

func fileExists(p string) (ok bool, err error) {
	_, err = os.Stat(p)
	switch {
	case err == nil:
		return true, nil
	case os.IsNotExist(err):
		return false, nil
	default:
		return false, err
	}
}

// getDirWork returns the dlsite.Work for the given directory.  Only
// the directory filename is used.
func getDirWork(f *dlsite.Fetcher, p string) (*dlsite.Work, error) {
	fn := filepath.Base(p)
	r := codes.ParseRJCode(fn)
	if r == "" {
		return nil, fmt.Errorf("invalid work filename %s", fn)
	}
	return f.FetchWork(codes.WorkCode(r))
}
