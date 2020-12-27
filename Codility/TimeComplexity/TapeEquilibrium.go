// A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

// Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

// The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

// In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

// For example, consider array A such that:
//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3

// We can split this tape in four places:

//         P = 1, difference = |3 − 10| = 7
//         P = 2, difference = |4 − 9| = 5
//         P = 3, difference = |6 − 7| = 1
//         P = 4, difference = |10 − 3| = 7

// Write a function:

//     func Solution(A []int) int

// that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

// For example, given:
//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3

// the function should return 1, as explained above.

// Write an efficient algorithm for the following assumptions:

//         N is an integer within the range [2..100,000];
//         each element of array A is an integer within the range [−1,000..1,000].

// 69 %
// func Solution(A []int) int {

// 	min, tmp := 0, 0
// 	P := len(A) - 1
// 	sum1, sum2 := 0, 0

// 	if len(A) == 2 {
// 		if A[0] > A[1] {
// 			return A[0] - A[1]
// 		} else if A[1] > A[0] {
// 			return A[1] - A[0]
// 		} else {
// 			return 0
// 		}
// 	}

// 	for i := 1; i <= P; i++ {
// 		for j := 0; j < i; j++ {
// 			sum1 += A[j]
// 		}
// 		for k := len(A); k > i; k-- {
// 			sum2 += A[k-1]
// 		}

// 		if sum1 > sum2 {
// 			tmp = sum1 - sum2
// 			if tmp < min || i == 1 {
// 				min = tmp
// 			}
// 		} else if sum2 > sum1 {
// 			tmp = sum2 - sum1
// 			if tmp < min || i == 1 {
// 				min = tmp
// 			}
// 		} else {
// 			return 0
// 		}
// 		tmp = 0
// 		sum1 = 0
// 		sum2 = 0
// 	}
// 	return min
// }

// 100 %
package solution

import "math"

func Solution(A []int) int {

	arraySum := 0
	currentMin := 1<<32 - 1

	for _, value := range A {
		arraySum += value
	}
	lhs := A[0]
	rhs := arraySum - lhs

	for i := 1; i < len(A); i++ {
		diff := int(math.Abs(float64(lhs) - float64(rhs)))

		if diff < currentMin {
			currentMin = diff
		}
		lhs += A[i]
		rhs -= A[i]
	}
	return currentMin
}
