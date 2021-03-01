package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}
	mag, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}
	switch {
	case mag >= 10:
		fmt.Printf("%g is massive.\n", mag)
	case mag >= 8:
		fmt.Printf("%g is great.\n", mag)
	case mag >= 7:
		fmt.Printf("%g is major.\n", mag)
	case mag >= 6:
		fmt.Printf("%g is strong.\n", mag)
	case mag >= 5:
		fmt.Printf("%g is moderate.\n", mag)
	case mag >= 4:
		fmt.Printf("%g is light.\n", mag)
	case mag >= 3:
		fmt.Printf("%g is minor.\n", mag)
	case mag >= 2:
		fmt.Printf("%g is very minor.\n", mag)
	default:
		fmt.Printf("%g is micro.\n", mag)
	}

}
