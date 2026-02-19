package main

import "fmt"

func printSlice[T int | string](items []T) {
	for _, i := range items {
		fmt.Println(i)
	}
}

type Stack[T int | string] struct {
	elements []T
}

func main() {

	stack := Stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println(stack)

	stackTwo := Stack[string]{
		elements: []string{"First", "Second"},
	}
	fmt.Println(stackTwo)

	printSlice([]int{1, 2, 3, 4, 5})
	printSlice([]string{"Anik", "ANik"})

}
