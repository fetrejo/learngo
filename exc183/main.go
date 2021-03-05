package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	words := filepath.SplitList(os.Getenv("PATH"))
	query := os.Args[1:]

queries:
	for _, q := range query {
		// case insensitive search
		q = strings.ToLower(q)

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
