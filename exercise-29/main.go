package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 42,
		"pedro": 26,
	}

	for k, v := range m {
		fmt.Printf("Key %v \t value %v\n", k, v)
	}

	age1 := m["james"]
	fmt.Println("The age of Bond", age1)
	if v, ok := m["james"]; ok {
		fmt.Println("There is a bond lookup entry, and bond's age is", v)
	}

	age2 := m["Q"]
	fmt.Println("The age of Q", age2)
	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q", v)
	}
}
