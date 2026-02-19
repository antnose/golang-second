package main

import "fmt"

func main() {
	// 06 if else
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("Adult")
	// }

	// if age := 14; age >= 10 {
	// 	fmt.Println("Per is an adult", age)
	// } else {
	// 	fmt.Println("Not adult")
	// }

	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("It's an integer")
	// 	case string:
	// 		fmt.Println("It's string")
	// 	case bool:
	// 		fmt.Println("Boolean", t)
	// 	}
	// }

	// whoAmI(12)

	// 08 array
	// var nums [4]int
	// nums[0] = 1

	// fmt.Println(nums)

	// numx := [3]int{1, 2, 3}
	// fmt.Println(numx)

	// nums := [2][2]int{{1, 4}, {5, 7}}
	// fmt.Println(nums)

	// 09 slices
	// var nums []int

	// fmt.Println(nums == nil)

	// fmt.Println(len(nums))

	// var nums = make([]int, 2)

	// nums = append(nums, 1)
	// fmt.Println(nums, "\n", len(nums), "\n", cap(nums))

	// nums = append(nums, 2)
	// fmt.Println(nums, "\n", len(nums), "\n", cap(nums))

	// nums = append(nums, 3)
	// fmt.Println(nums, "\n", len(nums), "\n", cap(nums))

	// 10 maps
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// fmt.Println(m["name"], m["area"])

	// fmt.Println(m["phone"])

	// m2 := map[string]int{"phone": 0}
	// fmt.Println(m2)

	// k, ok := m["name"]
	// if ok {
	// 	fmt.Println("all ok", ok)
	// } else {
	// 	fmt.Println("not ok", k)
	// }

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(v, k)
	}

}
