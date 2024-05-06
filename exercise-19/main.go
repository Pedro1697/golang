package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v \t", x)

	if x <= 100 {
		fmt.Printf("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Printf("between 101 and 200")
	} else {
		fmt.Printf("greater than 200")
	}
}
