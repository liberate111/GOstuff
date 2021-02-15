// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321

// Example 2:

// Input: x = -123
// Output: -321

// Example 3:

// Input: x = 120
// Output: 21

// Example 4:

// Input: x = 0
// Output: 0

// Constraints:

//     -231 <= x <= 231 - 1

// 1032 / 1032 test cases passed.
// Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.2 MB

func reverse(x int) int {
	min, max := -1<<31, 1<<31-1
	var result int

	for x != 0 {
		digit := x % 10
		result = result*10 + digit
		if result <= min || result >= max {
			return 0
		}
		x = x / 10
	}
	return result
}

// 1032 / 1032 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.2 MB

// import "strconv"

// func reverse(x int) int {
//     var strVal string
//     min, max := -1 << 31, 1 << 31 - 1
//     check := x

//     if x < 0 {
//         x *= -1
//     }

//     for _, v := range strconv.Itoa(x) {
//         strVal = string(v) + strVal
//     }

//     intVal, _ := strconv.Atoi(strVal)

//     if intVal <= min || intVal >= max {
//         return 0
//     }

//     if check < 0 {
//         intVal *= -1
//     }
//     return intVal
// }