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
}
