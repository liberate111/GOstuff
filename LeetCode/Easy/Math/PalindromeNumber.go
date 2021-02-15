// Given an integer x, return true if x is palindrome integer.

// An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.

// Example 1:

// Input: x = 121
// Output: true

// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Example 4:

// Input: x = -101
// Output: false

// Constraints:

//     -231 <= x <= 231 - 1

// Follow up: Could you solve it without converting the integer to a string?

// 11510 / 11510 test cases passed.
// Status: Accepted
// Runtime: 20 ms
// Memory Usage: 5.4 MB

func isPalindrome(x int) bool {
	x1 := x
	var pld int

	if x1 < 0 || x1%10 == 0 && x1 != 0 {
		return false
	}
	if x1 < 10 && x >= 0 {
		return true
	}

	for x1 != 0 {
		pld = (pld * 10) + (x1 % 10)
		x1 = x1 / 10
	}

	if pld == x {
		return true
	} else {
		return false
	}

}