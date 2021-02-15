// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Example 4:

// Input: s = ""
// Output: 0

// Constraints:

//     0 <= s.length <= 5 * 104
//     s consists of English letters, digits, symbols and spaces.

// 987 / 987 test cases passed.
// Status: Accepted
// Runtime: 4 ms
// Memory Usage: 4 MB

import "math"

func lengthOfLongestSubstring(s string) int {
	max, start := 0, 0
	m := make(map[byte]int, len(s))

	for i := range s {
		if idx, ok := m[s[i]]; ok && idx >= start {
			start = idx + 1
		}
		m[s[i]] = i
		max = int(math.Max(float64(max), float64(i-start+1)))
	}
	return max
}