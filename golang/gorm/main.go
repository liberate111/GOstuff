package main

import (
	"crypto/md5"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Email    string
	Username string
	Password string
}

func main() {
	db, err := gorm.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	user := User{Email: "test1@gmail.com", Username: "test1", Password: fmt.Sprintf("%x", md5.Sum([]byte("password")))}
	db.Save(&user)
	fmt.Println("after create")
	PrintUser(db)

	var users []User
	db.Find(&users, "email = ?", "test1@gmail.com").Update("email", "test2@gmail.com")
	fmt.Println("after update")
	PrintUser(db)

	db.Delete(&User{}, "username = ?", "test1")
	fmt.Println("after delete")
	PrintUser(db)
}

func PrintUser(db *gorm.DB) {
	var users []User
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.Email, user.Username, user.Password)
	}
}
