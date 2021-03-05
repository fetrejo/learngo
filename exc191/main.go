package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: [Username]")
		return
	}
	usr := os.Args[1]
	arr := [...]string{
		"good 👍",
		"bad 👎",
		"sad 😞",
		"happy 😀",
		"awesome 😎",
		"terrible 😩",
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(arr))
	fmt.Printf("%s feels %s\n", usr, arr[i])
}
