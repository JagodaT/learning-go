package main

import "fmt"

func main() {

	var array [10]int

	for i := 0; i < 10; i++ {
		array[i] = i + 1
	}
	fmt.Println(array)

}
