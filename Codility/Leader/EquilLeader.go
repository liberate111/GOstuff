//Find the index S such that the leaders of the sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N - 1] are the same.

// A non-empty array A consisting of N integers is given.

// The leader of this array is the value that occurs in more than half of the elements of A.

// An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

// For example, given array A such that:
//     A[0] = 4
//     A[1] = 3
//     A[2] = 4
//     A[3] = 4
//     A[4] = 4
//     A[5] = 2

// we can find two equi leaders:

//         0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
//         2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.

// The goal is to count the number of equi leaders.

// Write a function:

//     func Solution(A []int) int

// that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

// For example, given:
//     A[0] = 4
//     A[1] = 3
//     A[2] = 4
//     A[3] = 4
//     A[4] = 4
//     A[5] = 2

// the function should return 2, as explained above.

// Write an efficient algorithm for the following assumptions:

//         N is an integer within the range [1..100,000];
//         each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].

package solution

func Solution(A []int) int {
	elem := 1 << 31
	countElem := 0

	for _, v := range A {
		if elem == 1<<31 {
			elem = v
			countElem = 1
		} else {
			if v == elem {
				countElem++
			} else {
				countElem--
				if countElem == 0 {
					elem = 1 << 31
				}
			}
		}
	}

	if countElem == 0 {
		return 0
	}

	countAll := 0
	for _, v := range A {
		if v == elem {
			countAll++
		}
	}

	if countAll < int(len(A)/2) {
		return 0
	}

	countLeft := 0
	count := 0
	for i, v := range A {
		if v == elem {
			countLeft++
		}
		if countLeft > int((i+1)/2) {
			countRight := countAll - countLeft
			if countRight > int((len(A)-(i+1))/2) {
				count++
			}
		}
	}
	return count

}
