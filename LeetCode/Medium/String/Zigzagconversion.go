// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R

// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

// Constraints:

//     1 <= s.length <= 1000
//     s consists of English letters (lower-case and upper-case), ',' and '.'.
//     1 <= numRows <= 1000

// 1157 / 1157 test cases passed.
// Status: Accepted
// Runtime: 4 ms
// Memory Usage: 3.9 MB

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ret := make([]byte, 0, len(s))
	for d, n := 2*numRows-2, 0; n < numRows; n++ {
		for i := n; i < len(s); i += d {
			// Append normal row
			ret = append(ret, s[i])
			// Append Zigzag row
			if n != 0 && n != numRows-1 && i+d-2*n < len(s) {
				ret = append(ret, s[i+d-2*n])
			}
		}
	}
	return string(ret)
}