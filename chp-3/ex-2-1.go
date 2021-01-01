package main

import "fmt"

func bubbleSort(list []int) {
	var swapped bool = true
	for swapped {
		swapped = false
		for i := 1; i < len(list)-1; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				swapped = true
			}
		}
	}
}

func main() {

	list := []int{3, 5, 0, 2, 6, 4, 1, 32, 14}
	fmt.Println(list)

	bubbleSort(list)

	fmt.Println(list)

}
