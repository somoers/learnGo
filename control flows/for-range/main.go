package main

import "fmt"

func main() {
	x1 := []int{42, 43, 44, 45, 46, 47}
	for i, v := range x1 {
		fmt.Println("Going over slice", i, v)
	}

	m := map[string]int{
		"james": 42,
		"edson": 43,
	}
	for k, t := range m {
		fmt.Println("Going over map m", k, t)
	}
	fmt.Printf("The even numbers within 100 are ")
	for cn := 0; cn < 100; cn++ {
		if cn%2 == 0 {
			fmt.Printf("%v, ", cn)
		}

	}
}
