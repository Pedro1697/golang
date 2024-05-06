package main

import "fmt"

func main() {
	y := 0

	for {
		fmt.Println(y)
		if y > 20 {
			break
		}
		y++
	}
}
