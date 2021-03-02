package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong input.")
		return
	}
	n1, err1 := strconv.Atoi(os.Args[1])
	n2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("An error has ocurred.")
		return
	}
	if n1 >= n2 {
		fmt.Println("Second number has to be bigger.")
		return
	}
	var (
		i   = n1
		sum int
	)
	for {
		if i > n2 {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Printf("%d ", i)
		if i < n2-1 {
			fmt.Print(" + ")
		}
		sum += i
		i++
	}
	fmt.Printf("= %d\n", sum)
}
