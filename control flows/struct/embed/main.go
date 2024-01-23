package main

import "fmt"

type vehicle struct {
	engine string
	make   string
	model  string
	doors  int
	color  string
}

func main() {

	v1 := vehicle{
		engine: "vtti",
		make:   "Toyota",
		model:  "Vios",
		doors:  2,
		color:  "White",
	}
	fmt.Println(v1.engine)

}
