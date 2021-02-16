// Package main makes this package an executable program
package main

import (
	"fmt"
	"runtime"
)

/*
main func
Executes program
*/
func main() {
	fmt.Println(runtime.NumCPU())

}
