package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	path := os.Args[0]

	fmt.Printf("There are %d names.\n", count)
	fmt.Println(path)
	fmt.Println("Name is:", os.Args[1])
}
