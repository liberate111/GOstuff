// Determine whether a given string of parentheses (single type) is properly nested.
// A string S consisting of N characters is called properly nested if:

//         S is empty;
//         S has the form "(U)" where U is a properly nested string;
//         S has the form "VW" where V and W are properly nested strings.

// For example, string "(()(())())" is properly nested but string "())" isn't.

// Write a function:

//     func Solution(S string) int

// that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

// For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

// Write an efficient algorithm for the following assumptions:

//         N is an integer within the range [0..1,000,000];
//         string S consists only of the characters "(" and/or ")".

// 100 %
package solution

func Solution(S string) int {
	left := 0

	if len(S) == 0 {
		return 1
	}
	if len(S)%2 != 0 {
		return 0
	}

	for _, v := range S {
		val := string(v)
		if val == "(" {
			left += 1
		} else {
			if left == 0 {
				return 0
			} else {
				left -= 1
			}
		}
	}
	if left != 0 {
		return 0
	} else {
		return 1
	}
}
