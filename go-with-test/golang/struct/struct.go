package main

import (
	"fmt"
)

func main() {
	p1 := product{"Arduino", 100, 2}
	show(p1)
	update(&p1)
	show(p1)
	//fmt.Println(p1)

	var p2 product
	p2.name = "Rasberrypi"
	p2.price = 200
	p2.stock = 4
	show(p2)
	//fmt.Println(p2)
	//p1 = p1.clear()
	p1 = p1.setDiscount(1)
	show(p1)
}

type product struct {
	name  string
	price int
	stock int
}

func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}
func show(p product) {
	fmt.Println(p)
	fmt.Println(p.name)
}

func update(p *product) {
	p.price = p.price + 100
	p.stock = 10
}
