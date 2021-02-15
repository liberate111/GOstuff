// // CheatSheet
// // exact match
// abc      // matches "abc"

// // all character match
// .       // matches "a", "b", "x", "?", " "

// // letter match (A-Za-z0-9_)
// \w      // matches "a", "b", "x"

// // multiple letter match
// \w\w\w  // matches 3 letter
// \w+     // matches all possible consecutive letters
// \w{3}   // matches exact 3 consecutive letters
// \w{2,3} // 2 to 3
// \w*     // 0 or more

// [abc]   // match in abc
// r[ua]n  // "run" or "ran"

// // digits
// \d      // single digit
// \d\d    // 2 digit
// \d+     // as many as possible
// \d{4}   // just 4 digits
// \d{3,4} // 3 to 4 digit
// \d*     // 0 or more

// // whitespaces
// \t      // matches tab

// // groups
// (\d\w\d) // matches 1a1, 2e3 etc

package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match(`\d`, []byte("1 12 23")))   // true <nil>
	fmt.Println(regexp.Match(`\d\d`, []byte("1 12 23"))) // true <nil>
	fmt.Println(regexp.Match(`\w+`, []byte("1 12 23")))  // true <nil>

	re, e := regexp.Compile(`\w{3}`)
	CheckError(e)

	fmt.Println(string(re.Find([]byte("abc aaaa abcdef aaa abc adc")))) // abc
	fmt.Printf("%q\n", re.FindAll([]byte("abc aaaa abcdef aaa abc adc"), -1))

	re = regexp.MustCompile(`r[ua]n`)
	fmt.Printf("%q\n", re.FindAll([]byte("ran run run abc def abcd"), -1)) // ["ran" "run" "run"]

	re, e = regexp.Compile(`r[ua]n`)
	CheckError(e)
	fmt.Printf("%q\n", re.FindAll([]byte("ran run run abc def abcd"), -1)) // ["ran" "run" "run"]

	s := regexp.MustCompile(`a`).Split("abababacafagadatarat", 7) // [ b b b c f gadatarat]
	fmt.Println(s)
	fmt.Printf("%q\n", s)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
