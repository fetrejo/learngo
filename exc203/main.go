package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	var sep placeholder

	for {
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		if sec%2 == 0 {
			sep = colon
		} else {
			sep = blank
		}

		// [8][5]string
		clock := [...]placeholder{
			// extract the digits: 17 becomes, 1 and 7 respectively
			digits[hour/10], digits[hour%10],
			sep,
			digits[min/10], digits[min%10],
			sep,
			digits[sec/10], digits[sec%10],
		}

		screen.Clear()
		screen.MoveTopLeft()
		if sec%10 == 0 {
			for line := range alarm[0] {
				for letter := range alarm {
					fmt.Print(alarm[letter][line], "  ")
				}
				fmt.Println()
			}
		} else {
			for line := range clock[0] {
				for digit := range clock {
					fmt.Print(clock[digit][line], "  ")
				}
				fmt.Println()
			}
		}
		time.Sleep(time.Second)
	}
}
