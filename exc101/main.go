package main

//many exercises in one
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		width  uint8
		height uint16
	)

	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	fmt.Println("are they equal?", uint16(width) == height)

	val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("int8 value is :", int8(val))

	// 2nd argument is an int16
	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Println("int16 value is:", int16(val))

	// 3rd argument is an int32
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val))

	// 4th argument is an int64
	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Println("int64 value is:", int64(val))

	// 5th argument is a number in bits. And its int8.
	// ParseInt(.., 2, ...) -> 2 means binary base
	val, _ = strconv.ParseInt(os.Args[5], 2, 8)
	fmt.Printf("%s is: %d\n", os.Args[5], int8(val))

	t, _ := time.ParseDuration("1h30m")
	t *= time.Duration(val)

	type Feet float64
	type Metres float64

	var (
		feet   Feet
		meters Metres
		va     float64
	)

	va, _ = strconv.ParseFloat(os.Args[1], 64)
	feet = Feet(va)

	meters = Metres(feet * 0.3048)

	fmt.Printf("%g feet is %g metres.\n", feet, meters)

	fmt.Println(t)

	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(celsiusDegree) * Celsius(factor)
	fahr *= Fahrenheit(fahrDegree) * Fahrenheit(factor)

	fmt.Println(celsius, fahr)
}
