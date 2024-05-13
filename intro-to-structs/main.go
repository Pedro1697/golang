package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Pedor",
		last:  "Aguilar",
		age:   32,
	}

	fmt.Println(p1)

}
