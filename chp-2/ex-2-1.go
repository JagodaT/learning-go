package main

import "fmt"

func main() {

	floatArray := []float64{1, 2.5, 3, 55}
	var temp float64

	for _, v := range floatArray {
		temp += v
	}

	mean := temp / float64(len(floatArray))
	fmt.Println(mean)

}
