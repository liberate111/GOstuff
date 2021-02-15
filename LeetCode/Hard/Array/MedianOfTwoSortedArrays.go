// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// Follow up: The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Example 3:

// Input: nums1 = [0,0], nums2 = [0,0]
// Output: 0.00000

// Example 4:

// Input: nums1 = [], nums2 = [1]
// Output: 1.00000

// Example 5:

// Input: nums1 = [2], nums2 = []
// Output: 2.00000

// Constraints:

//     nums1.length == m
//     nums2.length == n
//     0 <= m <= 1000
//     0 <= n <= 1000
//     1 <= m + n <= 2000
//     -106 <= nums1[i], nums2[i] <= 106

// 2091 / 2091 test cases passed.
// Status: Accepted
// Runtime: 12 ms
// Memory Usage: 6 MB

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var med float64

	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	num := append(nums1, nums2...)
	sort.Ints(num)

	length := len(num)

	if length%2 == 0 {
		med = float64((num[length/2] + num[(length/2)-1])) / 2.00
	} else {
		med = float64(num[length/2])
	}
	return med
}