package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "yes",
		model: "v8",
		doors: 4,
		color: "black",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "yes",
		model: "v8",
		doors: 4,
		color: "black",
	}

	fmt.Println(v1)
	fmt.Println(v2)

}
