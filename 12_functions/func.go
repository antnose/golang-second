package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLang() (string, string, string) {
	return "golang", "Javascript", "c"
}

func main() {
	l1, l2, l3 := getLang()
	fmt.Println(l1, l2, l3)

	fmt.Println(add(12, 40))
}
