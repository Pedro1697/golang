package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {

	p1 := person{
		first:   "Pedro",
		last:    "Aguilar",
		flavors: []string{"chocolate", "uva"},
	}

	p2 := person{
		first:   "Pedro2",
		last:    "Aguilar2",
		flavors: []string{"chocolate", "uva"},
	}

	for _, v := range p1.flavors {
		fmt.Println(p1.first, v)

	}

	for _, v := range p2.flavors {
		fmt.Println(p2.first, v)

	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.flavors {
			fmt.Println(v.first, v.last, v2)
		}
	}

}
