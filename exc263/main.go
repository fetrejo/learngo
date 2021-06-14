package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	var null *computer

	if null == nil {
		fmt.Println("null computer is nill")
	}

	apple := &computer{brand: "apple"}

	newApple := apple

	if apple == newApple {
		fmt.Printf("apples are equal          : apple: %p newApple: %p\n",
			apple, newApple)
	}

	sony := &computer{brand: "sony"}

	if apple != sony {
		fmt.Printf("apple and sony are not equal          : apple: %p sony: %p\n",
			apple, sony)
	}

	appleVal := *apple

	fmt.Printf("apple                     : %p %p\n", &apple, apple)
	fmt.Printf("appleVal                  : %p\n", &appleVal)

	if *apple == appleVal {
		fmt.Println("apple and appleVal are equal")

		fmt.Printf("apple                     : %+v — appleVal: %+v\n",
			*apple, appleVal)
	}

	change(sony, "hp")
	fmt.Printf("sony's new brand: %s\n", sony.brand)

	fmt.Printf("appleVal                  : %+v\n", valueOf(apple))

	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}

func change(c *computer, brand string) {
	c.brand = brand
}

func valueOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
