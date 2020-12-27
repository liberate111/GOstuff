package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnWhileBreak()
}

func fnFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("IndexfnFor %d\n", i)
	}
}

func fnWhile() {
	i := 0
	for i < 5 {
		fmt.Printf("IndexfnWhile %d\n", i)
		i++
	}
}

func fnWhileBreak() {
	i := 0
	for true {
		if i > 5 {
			break
		}
		fmt.Printf("IndexBreak %d\n", i)
		i++
	}
}
