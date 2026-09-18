package main

import (
	"fmt"
	"io"
	"os"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	var r io.Reader

	if len(os.Args) > 1 {
		var readers []io.Reader
		for _, path := range os.Args[1:] {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
			defer f.Close()
			readers = append(readers, f)
		}
		r = io.MultiReader(readers...)
	} else {
		r = os.Stdin
	}

	counts, err := countWords(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}

	result := sortedCount(counts)
	printReport(os.Stdout, result)

}
