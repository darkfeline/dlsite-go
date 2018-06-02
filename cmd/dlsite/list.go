package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/dsutil"
	"go.felesatra.moe/subcommands"
)

func init() {
	commands = append(commands, subcommands.New("list", listCmd))
}

func listCmd(args []string) {
	f := flag.NewFlagSet("list", flag.ExitOnError)
	var i bool
	f.BoolVar(&i, "info", false, "Fetch work info also")
	f.Parse(args[1:])
	if err := listMain(i); err != nil {
		log.Fatal(err)
	}
}

func listMain(info bool) error {
	if !info {
		return mapCodes(os.Stdin, func(r dlsite.RJCode) error {
			fmt.Println(r)
			return nil
		})
	}
	c := dsutil.DefaultCache()
	defer c.Close()
	return mapCodes(os.Stdin, func(r dlsite.RJCode) error {
		w, err := dsutil.Fetch(c, r)
		if err != nil {
			return errors.Wrap(err, "fetch work")
		}
		printWork(os.Stdout, w)
		os.Stdout.Write([]byte("\n"))
		return nil
	})
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
