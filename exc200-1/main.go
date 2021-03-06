package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}
	n, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Wrong value. Error.")
		return
	}
	const (
		EUR = iota
		GBP
		JPY
	)
	arr := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}
	fmt.Printf("%.2f USD is %.2f EUR\n", n, arr[EUR]*n)
	fmt.Printf("%.2f USD is %.2f GBP\n", n, arr[GBP]*n)
	fmt.Printf("%.2f USD is %.2f JPY\n", n, arr[JPY]*n)
}
