package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4}
	var arr2 [3]string
	arr2[0], arr2[1], arr2[2] = "an", "ios", "react"

	for _, v := range arr1 {
		fmt.Println(v)
	}

	for _, v := range arr2 {
		fmt.Println(v)
	}
	fmt.Println(arr1[0])

	arr3 := [4]int{}
	fmt.Println(arr3)
}
