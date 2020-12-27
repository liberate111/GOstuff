package solution

func Solution(A []int, K int) []int {
	length := len(A)
	if length > 0 {
		if K > length {
			K = K % length
		}
		A = append(A[length-K:length], A[0:length-K]...)
	}
	return A
}

// func Solution(A []int, K int) []int {
// 	// write your code in Go 1.4
// 	B := make([]int, 5)
// 	check := A[0]
// 	count := 0

// 	for _, v := range A {
// 		if check == v {
// 			count++
// 		}
// 	}
// 	if count == len(A) {
// 		return A
// 	}

// 	if K == 0 || len(A)%K == 0 || len(A) == 0 {
// 		return A
// 	} else {
// 		s := K % len(A)
// 		//fmt.Println(s)
// 		for s > 0 {
// 			for i, _ := range B {
// 				if i == 0 {
// 					B[i] = A[len(A)-1]
// 				} else {
// 					B[i] = A[i-1]
// 					//fmt.Println(A[i - 1])
// 				}
// 			}
// 			//fmt.Println("B:",B)
// 			for i, _ := range A {
// 				A[i] = B[i]
// 			}
// 			//fmt.Println("A:",A)
// 			s--
// 		}
// 		return B
// 	}
// }

/*
An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).
The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.
Write a function:
    func Solution(A []int, K int) []int
that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.
For example, given
    A = [3, 8, 9, 7, 6]
    K = 3
the function should return [9, 7, 6, 3, 8]. Three rotations were made:
    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
For another example, given
    A = [0, 0, 0]
    K = 1
the function should return [0, 0, 0]
Given
    A = [1, 2, 3, 4]
    K = 4
the function should return [1, 2, 3, 4]
Assume that:
        N and K are integers within the range [0..100];
        each element of array A is an integer within the range [−1,000..1,000].
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
*/
