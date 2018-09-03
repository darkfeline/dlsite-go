package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/subcommands"
	"github.com/pkg/errors"

	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/cache"
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
			fmt.Fprintf(os.Stderr, "Could not get current directory: %s", err)
			return subcommands.ExitFailure
		}
	case 1:
		dir = f.Arg(0)
	default:
		fmt.Fprint(os.Stderr, c.Usage())
		return subcommands.ExitUsageError
	}
	if err := orgMain(dir, c.dry, c.all, c.desc); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func orgMain(dir string, dry, all, desc bool) error {
	var w []wPath
	var err error
	if all {
		w, err = findWorksAll(dir)
	} else {
		w, err = findWorks(dir)
	}
	if err != nil {
		return errors.Wrap(err, "find works")
	}
	c, err := cache.OpenDefault()
	if err != nil {
		return errors.Wrap(err, "open cache")
	}
	defer c.Close()
	for _, w := range w {
		log.Printf("Organizing %s", w)
		err := organizeWork(c, dir, w, dry, desc)
		if err != nil {
			return err
		}
	}
	if !dry {
		err := removeEmptyDirs(dir)
		if err != nil {
			log.Printf("Error removing empty dirs: %s", err)
		}
	}
	return nil
}

type wPath string

func findWorks(dir string) ([]wPath, error) {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var w []wPath
	for _, fi := range fi {
		if dlsite.Parse(fi.Name()) != "" {
			w = append(w, wPath(fi.Name()))
		}
	}
	return w, nil
}

// findWorksAll returns a slice of string paths of works.  The
// directory is searched recursively for works, which are returned by
// path relative to the given directory.
func findWorksAll(dir string) ([]wPath, error) {
	var w []wPath
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking %s: %s", path, err)
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		if dlsite.Parse(info.Name()) != "" {
			p, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			w = append(w, wPath(p))
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "walk")
	}
	return w, nil
}

func organizeWork(c *cache.Cache, topdir string, p wPath, dry, desc bool) error {
	w, err := pathWork(c, string(p))
	if err != nil {
		return errors.Wrap(err, "get work info")
	}
	new := workPath(w)
	if dry {
		if new != p {
			log.Printf("Would rename %s to %s", p, new)
		}
		return nil
	}
	if new != p {
		log.Printf("Moving %s", p)
		if err := renameWork(topdir, p, new); err != nil {
			return errors.Wrap(err, "rename work")
		}
		p = new
	}
	if desc {
		log.Printf("Adding description files for %s", p)
		if err := addDLSiteFiles(w, filepath.Join(topdir, string(p))); err != nil {
			return errors.Wrap(err, "add desc files")
		}
	}
	return nil
}

func renameWork(top string, old, new wPath) error {
	oldp := filepath.Join(top, string(old))
	newp := filepath.Join(top, string(new))
	if err := os.MkdirAll(filepath.Dir(newp), 0777); err != nil {
		return err
	}
	if err := os.Rename(oldp, newp); err != nil {
		return err
	}
	return nil
}

const (
	descFile  = "dlsite-description.txt"
	trackFile = "dlsite-tracklist.txt"
)

func addDLSiteFiles(w *dlsite.Work, p string) error {
	if w.Description != "" {
		fp := filepath.Join(p, descFile)
		_, err := os.Stat(fp)
		if err == nil {
			return nil
		}
		if !os.IsNotExist(err) {
			return err
		}
		if err := ioutil.WriteFile(fp, []byte(w.Description), 0666); err != nil {
			return err
		}
	}
	if len(w.TrackList) != 0 {
		fp := filepath.Join(p, trackFile)
		_, err := os.Stat(fp)
		if err == nil {
			return nil
		}
		if !os.IsNotExist(err) {
			return err
		}
		var b strings.Builder
		for i, t := range w.TrackList {
			b.WriteString(fmt.Sprintf("%d. %s %s\n", i, t.Name, t.Text))
		}
		if err := ioutil.WriteFile(fp, []byte(b.String()), 0666); err != nil {
			return err
		}
	}
	return nil
}

func pathWork(c *cache.Cache, p string) (*dlsite.Work, error) {
	fn := filepath.Base(p)
	r := dlsite.Parse(fn)
	if r == "" {
		return nil, errors.Errorf("invalid work filename %s", fn)
	}
	w, err := c.Get(r)
	if err == nil {
		return w, nil
	}
	log.Printf("Could not get %s from cache: %s", r, err)
	w, err = dlsite.Fetch(r)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func workPath(w *dlsite.Work) wPath {
	return wPath(filepath.Join(escapeFilename(w.Maker), escapeFilename(w.Series),
		escapeFilename(fmt.Sprintf("%s %s", w.RJCode, w.Name))))
}

func removeEmptyDirs(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking to %s: %s", path, err)
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		f, err := ioutil.ReadDir(path)
		if err != nil {
			log.Print(err)
			return nil
		}
		if len(f) == 0 {
			if err := os.Remove(path); err != nil {
				log.Print(err)
			}
		}
		return nil
	})
}
