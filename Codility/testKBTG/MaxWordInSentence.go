package testKBTG

import (
	"fmt"
	"strings"
)

// func main() {
// 	fmt.Println(FindMaxWord("This is testCoder, Give me a job?"))
// }

func FindMaxWord(S string) int {
	max := 0
	s := strings.ReplaceAll(S, "?", ",")
	fmt.Println("#1 s: ", s)

	s = strings.ReplaceAll(s, "!", ",")
	fmt.Println("#2 s: ", s)

	sentence := strings.Split(s, ",") //sentence is []string
	fmt.Println("sentence: ", sentence)

	for _, v := range sentence {
		sen := strings.Fields(v) //sen is []string
		fmt.Println("sen: ", sen)

		word := len(sen)
		fmt.Println("word: ", word)

		if word > max {
			max = word
		}
	}
	return max
}
