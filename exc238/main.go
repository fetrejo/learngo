package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please input some items")
		return
	}
	sort.Strings(args)

	var data []byte
	for i, s := range args {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, s...)
		data = append(data, '\n')
	}

	err := ioutil.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
