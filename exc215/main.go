package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("provide a few numbers")
	}
	arr := os.Args[1:]
	var narr []int
	for _, v := range arr {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		} else {
			narr = append(narr, n)
		}
	}
	sort.Ints(narr)
	fmt.Printf("%d\n", narr)
}
