package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	// blank identifier to not use a returned value
	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("-------------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	fmt.Println(len(xs))

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
	{
		// Append in a Slice
		xi := []int{42, 43, 44}
		fmt.Println(xi)
		fmt.Println("-------------")
		// variadic parameter
		xi = append(xi, 45, 46, 47, 99, 777)
		fmt.Println(xi)
		fmt.Println("-------------")
	}
	{
		// Slicing a slice
		xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Printf("xi - %#v\n", xi)
		fmt.Println("-------------")

		// [inclusive:exclusive]
		fmt.Printf("xi - %#v\n", xi[0:4])
		fmt.Println("-------------")

		// [:exclusive]
		fmt.Printf("xi - %#v\n", xi[:7])
		fmt.Println("-------------")

		// [inclusive:]
		fmt.Printf("xi - %#v\n", xi[4:])
		fmt.Println("-------------")

		// [:]
		fmt.Printf("xi - %#v\n", xi[:])
		fmt.Println("-------------")
	}
	{
		//Deleting from a slice
		xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Printf("xi - %#v\n", xi)
		fmt.Println("-------------")

		xi = append(xi[:4], xi[5:]...)
		fmt.Printf("xi - %#v\n", xi)
		fmt.Println("-------------")
	}
	{
		// func make
		xi := make([]int, 0, 10)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
		xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println(xi)
		fmt.Println("------------")
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
		fmt.Println("------------")
		xi = append(xi, 10, 11, 12, 13)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))

		x2 := make([]int, 2, 5)
		fmt.Println(x2)
		fmt.Println(len(x2))
		fmt.Println(cap(x2))
	}
}
