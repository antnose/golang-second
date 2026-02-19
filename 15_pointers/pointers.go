package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In change num", num)
// }

// By reference
func chNum(num *int) {
	*num = 5
	fmt.Println("In change num", *num)

}

func main() {
	num := 1

	chNum(&num)

	fmt.Println("After change", num)

	// changeNum(num)

	// fmt.Println("After change num in main func", num)
}
