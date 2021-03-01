package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a valid input.\n")
		return
	}
	age, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("An error has occurred.\n")
		return
	} else if age < 0 {
		fmt.Println("Please provide a valid input.\n")
		return
	}

	if age > 17 {
		fmt.Println("R-Rated\n")
	} else if (13 <= age) && (age <= 17) {
		fmt.Println("PG-13\n")
	} else {
		fmt.Println("PG-Rated\n")
	}

}
