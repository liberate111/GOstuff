package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	//Variables
	var name = "John"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	var username string = "John123"
	fmt.Println(username)
	var numberOfUser int
	fmt.Println(numberOfUser)

	number := 4.00
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))

	i, j, k := 1, 2, 3
	fmt.Println(i, "+", j, "=", i+j, "or", k)

	//Convert Types
	strVar := "100"
	intVar, _ := strconv.Atoi(strVar)
	fmt.Println(intVar)

	strVar1 := "-52541"
	if intVar1, err := strconv.ParseInt(strVar1, 10, 32); err == nil {
		fmt.Println(intVar1)
	}

	strVar2 := "101010101010101010"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
	fmt.Println(intVar2)

	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Printf("%T, %v\n", f, f)

	s1 := "-3.141"
	f1, _ := strconv.ParseFloat(s1, 8)
	fmt.Printf("%T, %v\n", f1, f1)

	s2 := "-3.141"
	f2, _ := strconv.ParseFloat(s2, 32)
	fmt.Printf("%T, %v\n", f2, f2)

	//Command Line arguments
	programName := os.Args[0]
	fmt.Println(programName)

	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d", argLength)

	for i, a := range os.Args[1:] {
		fmt.Printf("Arg %d is %s\n", i, a)
	}

	//flag
	//declaration
	var uname string
	var pass string

	//flag declaration
	flag.StringVar(&uname, "u", "root", "username. Default is root")
	flag.StringVar(&pass, "p", "password", "password. Default is password")

	//after declaring flags we need to call it

	flag.Usage = func() {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("./variable -u username -p password\n")
	}
	flag.Parse()
	//check if cli params match
	if uname == "root" && pass == "password" {
		fmt.Println("Login")
	} else {
		fmt.Println("Invalid")
	}

	//go run .\variable.go -u root -p password

}
