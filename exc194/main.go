package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...]string{"Hans", "Peter", "John"} // The names of your best three friends
	distances := [...]int{1, 2, 3, 4, 5}          // The distances to five different locations
	data := [...]byte{1, 2, 3, 4, 5}              // A data buffer with five bytes of capacity
	ratios := [...]float64{4.76}                  // Currency exchange ratios only for a single currency
	alives := [...]bool{true, true, false, true}  // Up/Down status of four different web servers
	var zero []byte                               // A byte array that doesn't occupy memory space

	// =========================================================================

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	// =========================================================================

	// you know how this works :) don't be freaked out!
	fmt.Printf(`
%s
FOR RANGES
%[1]s
`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}
