package main

import (
	"fmt"
)

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {
	seq := initSeq()
	fmt.Println("", seq())
	fmt.Println("", seq())

	fmt.Println(initSeq()())
	fmt.Println(initSeq()())
}
