package main

import "fmt"

var x = 40

const y = 50

func main() {

	if x == 42 {
		fmt.Println("Equal")
	} else if x < 42 {
		fmt.Println("Less")
	} else {
		fmt.Println("wala")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("The value of i is %v\n", i)
	}

	for x := 0; x < 20; x++ {
		if x%2 != 0 {
			continue
		}
		fmt.Printf("Couting even number %v\n", x)
	}

	for k := 0; k < 5; k++ {
		fmt.Println("---")
		for l := 0; l < 5; l++ {
			fmt.Printf("outer loop %v \t\t inner loop %v\n", k, l)
		}
	}

}
