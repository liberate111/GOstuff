package employee

import (
	"fmt"
)

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(
	firstname string,
	lastname string,
	totalleaves int,
	leavestaken int) employee {
	e := employee{
		FirstName:   firstname,
		LastName:    lastname,
		TotalLeaves: totalleaves,
		LeavesTaken: leavestaken}

	return e
}
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
