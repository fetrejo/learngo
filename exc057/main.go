package main

import (
	"fmt"
	"path"
)

func main() {
	_, b := multi()
	color, color2 := "red", "blue"
	color, color2 = color2, color
	dir, _ := path.Split("secret/file.txt")
	fmt.Println(b, color, color2, dir)
}

func multi() (int, int) {
	return 5, 4
}
