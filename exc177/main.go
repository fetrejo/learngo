package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 3 {
		// fmt.Println("Pick a number.")
		fmt.Printf(usage, maxTurns)
		return
	}
	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess1, err1 := strconv.Atoi(args[1])
	guess2, err2 := strconv.Atoi(args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}
	var sw int
	if guess1 > guess2 {
		sw = guess1
	} else {
		sw = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(sw + 1)
		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess1 || n == guess2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
