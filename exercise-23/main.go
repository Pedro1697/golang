package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		if x <= 3 {
			fmt.Printf("x has a value of %v\t\n", x)
		}

	}

}
