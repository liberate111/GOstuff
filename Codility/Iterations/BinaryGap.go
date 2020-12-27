package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution(32))
	fmt.Println(Solution(15))
	fmt.Println(Solution(1041))
	fmt.Println(Solution(9))
	fmt.Println(Solution(20))
	fmt.Println(Solution(2))
}

// !100 %
// func Solution(N int) int {

// 	//convert N to binary
// 	b := make([]int, 0)
// 	n := N
// 	check := true
// 	for check {
// 		s := n % 2
// 		b = append(b, s)
// 		n = n / 2
// 		if n == 1 {
// 			b = append(b, 1)
// 			check = false
// 		}
// 	}
// 	fmt.Println(b)
// 	cal := false
// 	gap := 0
// 	temp := 0

// 	//find binary 0 gap between 1
// 	for _, v := range b {
// 		if v == 1 {
// 			if !cal {
// 				cal = true
// 			}
// 			if temp > gap {
// 				gap = temp
// 			}
// 			temp = 0

// 		} else if cal && v == 0 {
// 			temp++
// 		}
// 	}
// 	return gap
// }

// A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

// For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

// 100 %
func Solution(N int) int {
	var result int
	var result_temp int
	var calc bool

	for N > 0 {
		if N%2 == 1 {
			if !calc {
				calc = true
			} else {
				if result_temp > result {
					result = result_temp
				}
				result_temp = 0
			}
		} else {
			if calc {
				result_temp++
			}
		}
		N = N / 2
	}

	return result
}
