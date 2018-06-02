package main

import (
	"fmt"
	"io"

	"go.felesatra.moe/dlsite"
)

func printWork(f io.Writer, w *dlsite.Work) (int, error) {
	const t = `%s
Name %s
Maker %s
Series %s
`
	return fmt.Fprintf(f, t, w.RJCode, w.Name, w.Maker, w.Series)
}
