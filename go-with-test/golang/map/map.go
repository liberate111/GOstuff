package main

import (
	"fmt"
)

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers["one"])

	colors := make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println(colors)
	fmt.Println(colors["green"])

	courses := make(map[string]map[string]int)

	courses["Android"] = map[string]int{"price": 200}

	courses["iOS"] = make(map[string]int)
	courses["iOS"]["price"] = 100
	courses["iOS"]["code"] = 101

	fmt.Println(courses)
}
