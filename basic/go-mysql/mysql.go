package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Name string
	Text string
}

func main() {
	db, e := sql.Open("mysql", "fufuu:fufuu@/posts")
	ErrorCheck(e)

	defer db.Close()
	PingDB(db)

	// Insert into DB
	/*
		stmt, e := db.Prepare("insert into posts(id, Name, Text) values (?, ?, ?)")
		ErrorCheck(e)
		defer stmt.Close()

		res, e := stmt.Exec("5", "post 5", "content 5")
		ErrorCheck(e)

		id, e := res.LastInsertId()
		ErrorCheck(e)

		fmt.Println("Insert id", id)
	*/
	//Insert id 0

	// Update data
	/*
		stmt, e := db.Prepare("update posts set Text=? where id=?")
		ErrorCheck(e)

		res, e := stmt.Exec("this is post 5", "5")
		ErrorCheck(e)

		row, e := res.RowsAffected()
		ErrorCheck(e)

		fmt.Println(row)
	*/

	// Query data
	rows, e := db.Query("select * from posts")
	ErrorCheck(e)

	var post = Post{}

	// iterate over rows
	for rows.Next() {
		e = rows.Scan(&post.Id, &post.Name, &post.Text)
		ErrorCheck(e)
		fmt.Println(post)
	}

	// Query data
	rows2, e := db.Query("select * from posts")
	ErrorCheck(e)

	var posts []Post
	var eachPost Post

	// iterate over rows
	for rows2.Next() {
		e = rows2.Scan(&eachPost.Id, &eachPost.Name, &eachPost.Text)
		ErrorCheck(e)

		posts = append(posts, eachPost)
		fmt.Println(eachPost)
	}
	fmt.Println(posts)
	fmt.Println(posts[0])

	// Delete date
	stmt, e := db.Prepare("delete from posts where id=?")
	ErrorCheck(e)

	res, e := stmt.Exec("5")
	ErrorCheck(e)

	row, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(row)
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
