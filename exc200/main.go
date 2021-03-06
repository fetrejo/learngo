package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [5][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}
	for i := range arr {
		n := arr[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}
}
