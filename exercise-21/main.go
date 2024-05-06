package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Print(x, y)

	if x < 4 && y < 4 {
		fmt.Print("both less than four")
	} else if x > 6 && y > 6 {
		fmt.Print("Both greater than 6")
	} else {
		fmt.Print("one is no equal")
	}
}
