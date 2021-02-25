package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		minsPerDay   = 1440
		weekDays     = 7
		hoursInDay   = 24
		hoursPerWeek = weekDays * hoursInDay
		weekCount    = 5
		home         = "Milky Way Galaxy"
		length       = len(home)
		pi           = 3.14159265358979323846264
		tau          = pi * 2
		width        = 25
		height       = width * 2
		later        = 10
	)
	const (
		Nov = 11 - iota
		Oct
		Sep
	)
	const (
		_ = iota
		Jan
		Feb
		Mar
	)
	const (
		Winter = 12 - iota
		_
		_
		Fall
		_
		_
		Summer
		_
		_
		Spring
	)
	fmt.Printf("Minutes in two weeks: %d\n", (minsPerDay*weekDays)*2)
	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*weekCount, weekCount)
	fmt.Printf("There are %d characters inside %q.\n", length, home)
	fmt.Printf("tau = %g\n", tau)
	fmt.Printf("area = %d\n", width*height)
	hours, _ := time.ParseDuration("1h")
	fmt.Printf("%s later...\n", hours*later)
	fmt.Println(Sep, Oct, Nov)
	fmt.Println(Jan, Feb, Mar)
	fmt.Println(Winter, Spring, Summer, Fall)
}
