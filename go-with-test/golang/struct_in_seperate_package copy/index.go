package main

import "structEmp/employee"

func main() {
	e := employee.Init(
		"Sam",
		"Smith",
		30,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"adssda",
		"asdsda",
		30,
		20)

	e.LeavesRemaining()
}

//Singleton
