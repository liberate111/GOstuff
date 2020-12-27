package main

import "fmt"

func main() {
	sample1 := []int{1, 3, 6, 4, 1, 2}
	result1 := Solution(sample1)
	fmt.Println(sample1, result1)

	sample2 := []int{1, 2, 3}
	result2 := Solution(sample2)
	fmt.Println(sample2, result2)

	sample3 := []int{-1, -2}
	result3 := Solution(sample3)
	fmt.Println(sample3, result3)

	sample4 := []int{1, 2, 3, 10, 1000, 9}
	result4 := Solution(sample4)
	fmt.Println(sample4, result4)
}

func Solution(A []int) int {

	m := make(map[int]int)

	s := 1
	b := true

	for _, v := range A {
		if v > 0 {
			m[v] = v
		}
	}

	for b {
		if _, ok := m[s]; !ok {
			b = false
		} else {
			s++
		}
	}
	return s
}

// This is a demo task.

// Write a function:

//     func Solution(A []int) int

// that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

// For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

// Given A = [1, 2, 3], the function should return 4.

// Given A = [−1, −3], the function should return 1.

// Write an efficient algorithm for the following assumptions:

//         N is an integer within the range [1..100,000];
//         each element of array A is an integer within the range [−1,000,000..1,000,000].
