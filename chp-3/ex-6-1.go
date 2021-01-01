package main

import "fmt"

func main() {

	p := plusTwo()
	fmt.Println(p(4))

}

func plusTwo() func(int) int {
	ret := func(a int) int {
		return a + 2
	}
	return ret
}
