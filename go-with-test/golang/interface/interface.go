package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rectangle":
		r := s.(rectangle)
		fmt.Printf("Reac width: %f height: %f\n", r.width, r.height)
	case "circle":
		c := s.(circle)
		fmt.Printf("Circle circle: %f \n", c.radius)
	}
}

func castToRectangle(s shape) {
	r, ok := s.(rectangle)
	if !ok {
		fmt.Println("Casting Error")
	} else {
		fmt.Printf("Casting Success w: %f h: %f\n", r.width, r.height)
	}
}

func main() {
	r := rectangle{10, 20}
	c := circle{2}

	fmt.Printf("Area : %f\n", r.area())
	fmt.Printf("Area : %f\n", c.area())

	showInfo(r)
	showInfo(c)

	castToRectangle(r)
	castToRectangle(c)
}
