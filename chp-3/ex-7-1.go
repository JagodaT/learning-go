package main

import "fmt"

func mapF(f func(int) int, x []int) []int {
	ret := make([]int, len(x))
	for i, v := range x {
		ret[i] = f(v)
	}
	return ret
}

func main() {

	testFunction := func(x int) int {
		return x * x
	}

	list := []int{1, 2, 3, 4, 5}

	fmt.Println(mapF(testFunction, list))

}
