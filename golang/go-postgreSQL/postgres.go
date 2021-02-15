package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "fufuu"
	dbname   = "postgres"
)

func main() {
	// connecting to PostgreSQL
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// Get data
	rows, err := db.Query(`SELECT "Name", "Roll" FROM "Students"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
		CheckError(err)

		fmt.Println(name, roll)
	}

	// Insert

	// hardcoded
	// insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// dynamic
	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	// _, e := db.Exec(insertDynStmt, "Bob4", 4)
	// CheckError(e)
	// _, e = db.Exec(insertDynStmt, "John", 2)
	// CheckError(e)

	// Delete
	// deleteStmt := `delete from "Students" where id=$1`
	// _, e := db.Exec(deleteStmt, 2)
	// CheckError(e)

	// Update
	// updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	// _, e := db.Exec(updateStmt, "Bob3", 3, 3)
	// CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
