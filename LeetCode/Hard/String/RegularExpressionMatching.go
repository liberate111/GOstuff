// Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where:

//     '.' Matches any single character.​​​​
//     '*' Matches zero or more of the preceding element.

// The matching should cover the entire input string (not partial).

// Example 1:

// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".

// Example 2:

// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

// Example 3:

// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".

// Example 4:

// Input: s = "aab", p = "c*a*b"
// Output: true
// Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".

// Example 5:

// Input: s = "mississippi", p = "mis*is*p*."
// Output: false

// Constraints:

//     0 <= s.length <= 20
//     0 <= p.length <= 30
//     s contains only lowercase English letters.
//     p contains only lowercase English letters, '.', and '*'.
//     It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.

// 447 / 447 test cases passed.
// Status: Accepted
// Runtime: 8 ms
// Memory Usage: 5.1 MB

func isMatch(s string, p string) bool {
	// ^ at beginning of text or line (flag m=true)
	// $ at end of text (like \z not Perl's \Z) or line (flag m=true)
	// . any character, possibly including newline (flag s=true)
	// x* zero or more x, prefer more
	return regexp.MustCompile("^" + p + "$").MatchString(s)
}