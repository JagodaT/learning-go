package main

func main() {

}

func popStack(stack []int) {

	idx := len(stack)
	if idx == 0 {
		return
	}

	stack = stack[:idx-2]

}

func pushStack(stack []int, k int) {
	stack = append(stack, k)
}
