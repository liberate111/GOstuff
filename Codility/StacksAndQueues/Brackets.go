//Determine whether a given string of parentheses (multiple types) is properly nested.
// A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

//         S is empty;
//         S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
//         S has the form "VW" where V and W are properly nested strings.

// For example, the string "{[()()]}" is properly nested but "([)()]" is not.

// Write a function:

//     func Solution(S string) int

// that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

// For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

// Write an efficient algorithm for the following assumptions:

//         N is an integer within the range [0..200,000];
//         string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".

// 100 %

package solution

func Solution(S string) int {
	arrLen := len(S)
	stack := make([]string, 0)
	pairs := make(map[string]string)
	pairs["{"] = "}"
	pairs["["] = "]"
	pairs["("] = ")"

	if arrLen == 0 {
		return 1
	}
	if arrLen%2 != 0 {
		return 0
	}

	for _, v := range S {
		val := string(v)
		if val == "{" || val == "[" || val == "(" {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return 0
			} else {
				last := stack[len(stack)-1]
				if val != pairs[string(last)] {
					return 0
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return 1
	} else {
		return 0
	}
}
