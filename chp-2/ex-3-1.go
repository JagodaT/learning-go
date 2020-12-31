package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		mod3, mod5 := i%3, i%5

		if i == 0 {
			fmt.Println(i)
			continue
		}

		if mod3 == 0 && mod5 == 0 {
			fmt.Println("FizzBuzz")
		} else if mod3 == 0 {
			fmt.Println("Fizz")
		} else if mod5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
