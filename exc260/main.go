package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("shakespeare.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)

	rx := regexp.MustCompile(`[^A-Za-z]+`)

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++

		word := rx.ReplaceAllString(in.Text(), "")
		word = strings.ToLower(word)
		words[word]++
	}

	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}
