// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

//     2 <= nums.length <= 103
//     -109 <= nums[i] <= 109
//     -109 <= target <= 109
//     Only one valid answer exists.

// Submission Detail
// 52 / 52 test cases passed.
// 	Status: Accepted
// Runtime: 4 ms
// Memory Usage: 3.2 MB

// func twoSum(nums []int, target int) []int {
//     checkSum := false
//     idx1 := 0
//     idx2 := 1
//     ans := make([]int, 0)
//     N := len(nums)
//     for !checkSum {
//         if nums[idx1] + nums[idx2] != target {
//             idx2++
//             if idx2 == N {
//                 idx1++
//                 idx2 = idx1 + 1
//                 if idx1 == N-1 {
//                     checkSum = true
//                     ans = append(ans, 0, 0)
//                     return ans
//                 }
//             }
//         } else {
//             ans = append(ans, idx1, idx2)
//             checkSum = true
//             return ans
//         }
//     }
//     return ans
// }

// 52 / 52 test cases passed.
// 	Status: Accepted
// Runtime: 4 ms
// Memory Usage: 3.2 MB

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}