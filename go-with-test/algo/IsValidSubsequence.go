package main

func IsValidSubsequence(array []int, sequence []int) bool {

	subIdx := 0
	for _, value := range array {
		if subIdx == len(sequence) {
			break
		}
		if sequence[subIdx] == value {
			subIdx++
		}
	}
	return subIdx == len(sequence)
}

//test
// package main

// import (
// 	"github.com/stretchr/testify/require"
// )

// func (s *TestSuite) TestCase1(t *TestCase) {
// 	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
// 	sequence := []int{1, 6, -1, 10}
// 	require.True(t, IsValidSubsequence(array, sequence))
// }
