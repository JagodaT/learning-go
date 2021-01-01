package main

import "fmt"

func main() {

	p := plusX(3)
	fmt.Println(p(4))

}

func plusX(x int) func(int) int {
	ret := func(a int) int {
		return x + a
	}
	return ret
}
