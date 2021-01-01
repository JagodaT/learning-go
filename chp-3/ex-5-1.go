package main

import "fmt"

func main() {
	printVarArgs(1, 4, 5, 2, 56, 3, 5)
}

func printVarArgs(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
	return
}
