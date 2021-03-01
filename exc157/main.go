package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() > 0 && t.Hour() < 4:
		fmt.Println("Good night.")
	case t.Hour() > 3 && t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() > 11 && t.Hour() < 8:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
