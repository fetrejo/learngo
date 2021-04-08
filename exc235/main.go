package main

import "fmt"

func main() {
	names := make([]string, 5)
	fmt.Printf("1st step %v\n", names)

	names = append(names, "einstein", "tesla", "aristo")
	fmt.Printf("2nd step %v\n", names)

	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristo")
	fmt.Printf("3rd step %v\n", names)

	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)]
	fmt.Printf("4th step %v\n", names)

	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	fmt.Printf("5th step (before append) %v\n", clone)
	clone = append(clone, names[0:2]...)
	fmt.Printf("5th step (after append) %v\n", clone)

	sliced := clone[1:4:4]
	sliced = append(sliced, "hypatia")
	clone[2] = "elder"
	fmt.Printf("6th step %v %v\n", clone, sliced)
}
