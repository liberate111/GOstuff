// Count factors of given number n.
// A positive integer D is a factor of a positive integer N if there exists an integer M such that N = D * M.

// For example, 6 is a factor of 24, because M = 4 satisfies the above condition (24 = 6 * 4).

// Write a function:

//     func Solution(N int) int

// that, given a positive integer N, returns the number of its factors.

// For example, given N = 24, the function should return 8, because 24 has 8 factors, namely 1, 2, 3, 4, 6, 8, 12, 24. There are no other factors of 24.

// Write an efficient algorithm for the following assumptions:

//         N is an integer within the range [1..2,147,483,647].

// 85 %
// 100 % correctness 66 % performance
// package solution

// // import "fmt"

// func Solution(N int) int {
// 	countFactor := 0
// 	lastFactor := 0

// 	for i := 1; i <= N; i++ {
// 		if i*lastFactor == N {
// 			countFactor *= 2
// 			return countFactor
// 		} else if i*i == N {
// 			countFactor = countFactor*2 + 1
// 			return countFactor
// 		} else if N%i == 0 {
// 			countFactor++
// 			lastFactor = i
// 		}
// 	}

// 	return countFactor
// }

// 100 %
package solution

func Solution(N int) int {
	i := 1
	count := 0

	for i*i <= N {
		if i*i == N {
			count++
		} else if N%i == 0 {
			count += 2
		}
		i++
	}
	return count
}
