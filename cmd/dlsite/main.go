package main

import (
	"fmt"
	"log"
	"os"

	"go.felesatra.moe/dlsite"
)

func main() {
	c := dlsite.Parse(os.Args[1])
	if c == "" {
		log.Printf("Invalid RJ code")
	}
	w, err := dlsite.Fetch(c)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", w)
}
