package main

import "fmt"

func main() {
	week := [4]string{"Monday", "Tuesday"}
	wend := [4]string{"Saturday", "Sunday"}
	fmt.Println(week != wend)
	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}
	fmt.Println(evens == evens2)

	image := [5]uint8{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
