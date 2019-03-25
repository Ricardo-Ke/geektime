package main

import "fmt"

func main() {
	// nums1 := []int{1, 2, 3, 4, 5, 6}
	// for i, v := range nums1 {
	// 	fmt.Println(i, v)
	// 	if i == 3 {
	// 		nums1[i] |= i
	// 	}
	// }
	// fmt.Println(nums1)

	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	case uint16:
		fmt.Println("ui8 or ui16")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Printf("unsupported type %T", t)
	}

	var a string
	fmt.Println(a == "")
}
