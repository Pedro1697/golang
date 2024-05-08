package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"d", "e", "f"}
	m["jenny"] = []string{"a", "b", "c"}
	m["doctor"] = []string{"g", "h", "i"}

	fmt.Printf("%#v", m)

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
