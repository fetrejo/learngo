package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	name := os.Args[1]
	msg := `hi ` + name +
		` how are you?`

	length := utf8.RuneCountInString(os.Args[1])

	fmt.Println(path, msg, length)
}
