package main

import "fmt"

func main() {

	fibonnacci(20)

}

func fibonnacci(n int) {

	var previous, current int = 0, 1
	if n == 0 {
		fmt.Println(n)
		return
	}
	for i := 0; i < n; i++ {
		fmt.Println(current)
		current, previous = current+previous, current
	}
	return
}
