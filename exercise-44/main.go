package main

import "fmt"

func main() {
	p1 := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "aguilar",
		friends: map[string]int{
			"q": 85,
			"a": 56,
		},
		favDrink: []string{"coca", "pepsi"},
	}

	for k, v := range p1.friends {
		fmt.Println(p1.first, k, v)
	}

	for _, v := range p1.favDrink {
		fmt.Println(p1.first, v)
	}
}
