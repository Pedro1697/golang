package main

import "fmt"

func main() {
	a := [5]int{}

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Print(a)
}
