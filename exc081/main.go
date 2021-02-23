package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	x := 5. / 2
	counter, factor := 45, 0.5

	var (
		radius = 10.
		area   float64
	)

	pi := math.Pi
	area = (radius * radius) * pi
	counter += 5
	factor -= 2

	rad, _ := strconv.ParseFloat(os.Args[1], 64)
	ar := math.Pow(rad, 2) * 4 * pi

	vol := math.Pow(rad, 3) * (4. / 3) * pi

	fmt.Println(x, float64(counter)*factor)
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
	fmt.Printf("radius: %g -> area: %.2f\n", rad, ar)
	fmt.Printf("radius: %g -> volume: %.2f\n", rad, vol)
}
