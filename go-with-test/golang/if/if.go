package main

import "fmt"

func main() {
	v := 11
	if v == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!=10")
	}

	if result := doSomething(); result == "ok" {
		fmt.Println("ok")
	} else {
		fmt.Println("!ok")
	}

	fnSwitchCase()
}

func doSomething() string {
	return "ok"
}

func fnSwitchCase() {
	i := 3
	switch i {
	case 0:
		fmt.Println("0")
		//break
	case 1:
		fmt.Println("1")
		//break
	default:
		fmt.Println("2")
		//break
	}
}
