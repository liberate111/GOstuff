package main

import "structEmp/employee"

func main() {
	e := employee.New(
		"Sam",
		"Smith",
		30,
		20)

	e.LeavesRemaining()
}
