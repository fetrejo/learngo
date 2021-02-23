package main

import (
	"fmt"
	"os"
)

func main() {
	a := 14
	s := "STR"
	temp := 17.4
	name, lastname := os.Args[1], os.Args[2]
	fmt.Printf("Test 1: %d\n", a)
	fmt.Printf("Test 2: %s\n", s)
	fmt.Printf("Test 3: %.1f\n", temp)
	fmt.Printf("Test 4: \"%s\"\n", s)
	fmt.Printf("Test 5: %T\n", a)
	fmt.Printf("Test 5: %s and %s\n", name, lastname)
}
