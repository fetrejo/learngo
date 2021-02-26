package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [username] [password]\n")
	} else if os.Args[1] == "jack" {
		if os.Args[2] == "1888" {
			fmt.Println("Access granted to jack.\n")
		} else {
			fmt.Println("Invalid password for jack.\n")
		}
	} else {
		fmt.Printf("Access denied for %q.\n", os.Args[1])
	}
}
