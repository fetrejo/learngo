package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non-numerics
			continue
		}

		if big.NewInt(int64(n)).ProbablyPrime(0) {
			fmt.Println(n, "is prime")
		} else {
			continue

		}
	}
}
