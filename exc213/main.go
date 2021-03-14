package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	nA := strings.Split(namesA, ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(nA)
	sort.Strings(namesB)

	if len(nA) == len(namesB) {
		for i := range nA {
			same := nA[i] == namesB[i]
			if same {
				continue
			} else {
				fmt.Println("Slices are different.")
				return
			}
		}
		fmt.Println("Slices are the same.")
	}
}
