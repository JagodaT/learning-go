package main

import "fmt"

func main() {

	i := 0
loop:
	fmt.Println(i)
	if i < 10 {
		i++
		goto loop
	}

}
