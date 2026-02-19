package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for i := range nums {
		total += i
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))

	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(nums...))
}
