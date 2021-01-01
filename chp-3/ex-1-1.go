package main

import "fmt"

func meanArray(floatArray []float64) float64 {

	var sum float64
	for _, v := range floatArray {
		sum += v
	}

	return sum / float64(len(floatArray))

}

func main() {

	array := []float64{1, 2, 3, 4, 5}

	fmt.Println(meanArray(array))

}
