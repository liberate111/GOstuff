package main

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	ID     int    `xml:"ID"`
	Title  string `xml:"Name"`
	Author string `xml:"Author"`
}

func main() {
	b := &Book{1, "Book1", "fufuu1"}
	b1 := Book{2, "Book2", "fufuu2"}

	fmt.Println(b)
	fmt.Println(b1)

	v, _ := xml.MarshalIndent(b, "", "   ")
	v1, _ := xml.MarshalIndent(b1, "", "   ")

	fmt.Println(string(v))
	fmt.Println(string(v1))

	x := `
		<Book>
			<ID>1</ID>
			<Name>Book3</Name>
			<Author>fufuu3</Author>
		</Book>
	`
	b3 := &Book{}
	xml.Unmarshal([]byte(x), b3)

	fmt.Println(b3)
	fmt.Printf("%+v\n", b3)

	xml.Unmarshal([]byte(v1), b3)
	fmt.Printf("%+v\n", b3)
}
