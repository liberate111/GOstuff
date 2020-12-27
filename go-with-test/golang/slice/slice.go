package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	var numbers2 []int

	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	showSlice(numbers)

	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	numbers2 = append(numbers2, 3)
	showSlice(numbers2)

	var numbers3 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	showSlice(numbers3)

	//remove heading
	numbers4 := numbers3[1:len(numbers3)]
	showSlice(numbers4)

	//remove tailing
	numbers4 = numbers3[0 : len(numbers3)-1]
	showSlice(numbers4)

	numbers3 = removeIndex(numbers3, 5)
	showSlice(numbers3)
}

func showSlice(s []int) {
	fmt.Printf("len= %d cap=%d slice%v\n", len(s), cap(s), s)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
