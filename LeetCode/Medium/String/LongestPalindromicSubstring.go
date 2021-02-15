// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.

// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Example 3:

// Input: s = "a"
// Output: "a"

// Example 4:

// Input: s = "ac"
// Output: "a"

// Constraints:

//     1 <= s.length <= 1000
//     s consist of only digits and English letters (lower-case and/or upper-case),

// 176 / 176 test cases passed.
// Status: Accepted
// Runtime: 4 ms
// Memory Usage: 2.7 MB

func longestPalindrome(s string) string {
	result := ""

	if len(s) == 0 {
		return result
	}
	if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		tmp := findPld(s, i, i)
		if len(tmp) > len(result) {
			result = tmp
		}

		tmp = findPld(s, i, i+1)
		if len(tmp) > len(result) {
			result = tmp
		}
	}
	return result
}

func findPld(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}