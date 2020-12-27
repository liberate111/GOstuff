package main

import "fmt"

func main() {
	course := []string{"an", "ios", "c"}

	for index, v := range course {
		fmt.Println(index+1, ".", v)
	}

	for _, v := range course {
		fmt.Println(v)
	}
}
