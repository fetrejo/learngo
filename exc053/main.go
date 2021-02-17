package main

import "fmt"

func main() {
	i := 314
	n, bol := 14, true
	sum := 27 + 3.5
	bol1, bol2 := true, true
	_ = bol2
	age, yourAge := 1, 2
	ratio, age := 0.5, 42
	fmt.Println(i, n, bol, sum, bol1, yourAge, ratio, age)
}
