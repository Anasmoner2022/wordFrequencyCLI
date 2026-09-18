package main

import (
	"fmt"
	"io"
)

func printReport(w io.Writer, output []WordCount) {
	for _, r := range output {
		fmt.Fprintf(w, "%s %d\n", r.Word, r.Count)
	}
}
