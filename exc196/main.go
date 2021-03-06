package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name := args[0]

	moods := [2][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	if args[1] == "positive" {
		fmt.Printf("%s feels %s\n", name, moods[0][n])
	} else {
		fmt.Printf("%s feels %s\n", name, moods[1][n])
	}
}
