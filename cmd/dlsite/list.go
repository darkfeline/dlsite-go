package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

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

	c := make(chan dlsite.RJCode)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if i {
			printWorks(c)
		} else {
			printCodes(c)
		}
		wg.Done()
	}()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		r := dlsite.Parse(s.Text())
		if r != "" {
			c <- r
		}
	}
	close(c)
	wg.Wait()
	if err := s.Err(); err != nil {
		log.Fatalf("Error reading lines: %s", err)
	}
}

func printCodes(c <-chan dlsite.RJCode) {
	for r := range c {
		fmt.Println(r)
	}
}

func printWorks(c <-chan dlsite.RJCode) {
	dc := dsutil.DefaultCache()
	defer dc.Close()
	for r := range c {
		w, err := dsutil.Fetch(dc, r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching work: %s\n", err)
		}
		printWork(os.Stdout, w)
		os.Stdout.Write([]byte("\n"))
	}
}
