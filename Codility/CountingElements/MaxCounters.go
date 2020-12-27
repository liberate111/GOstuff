// You are given N counters, initially set to 0, and you have two possible operations on them:

// increase(X) − counter X is increased by 1,
// max counter − all counters are set to the maximum value of any counter.

// A non-empty array A of M integers is given. This array represents consecutive operations:

// if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
// if A[K] = N + 1 then operation K is max counter.

// For example, given integer N = 5 and array A such that:
// A[0] = 3
// A[1] = 4
// A[2] = 4
// A[3] = 6
// A[4] = 1
// A[5] = 4
// A[6] = 4

// the values of the counters after each consecutive operation will be:
// (0, 0, 1, 0, 0)
// (0, 0, 1, 1, 0)
// (0, 0, 1, 2, 0)
// (2, 2, 2, 2, 2)
// (3, 2, 2, 2, 2)
// (3, 2, 2, 3, 2)
// (3, 2, 2, 4, 2)

// The goal is to calculate the value of every counter after all operations.

// Write a function:

// func Solution(N int, A []int) []int

// that, given an integer N and a non-empty array A consisting of M integers, returns a sequence of integers representing the values of the counters.

// Result array should be returned as an array of integers.

// For example, given:
// A[0] = 3
// A[1] = 4
// A[2] = 4
// A[3] = 6
// A[4] = 1
// A[5] = 4
// A[6] = 4

// the function should return [3, 2, 2, 4, 2], as explained above.

// Write an efficient algorithm for the following assumptions:

// N and M are integers within the range [1..100,000];
// each element of array A is an integer within the range [1..N + 1].

// 100 %
package solution

func Solution(N int, A []int) []int {

	result := make([]int, N)
	max := N + 1
	currentMax := 0
	base := 0

	// counter and adjust base
	// if no counter integer after count max, value not update to base
	for _, v := range A {
		if v == max {
			base = currentMax
		} else {
			if result[v-1] < base {
				result[v-1] = base
			}

			result[v-1]++
			if currentMax < result[v-1] {
				currentMax = result[v-1]
			}
		}
	}
	// update base that value don't count after count max
	for i, v := range result {
		if v < base {
			result[i] = base
		}
	}
	return result
}
