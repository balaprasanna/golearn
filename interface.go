package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	// L * B
	return r.width * r.height
}

func (r *rect) perimeter() float64 {
	// 2 * (L+B)
	return 2 * (r.width + r.height)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func shapeStarts(s shape) {
	fmt.Println("Area", s.area())
	fmt.Println("Perimeter", s.perimeter())
}

func main() {
	fmt.Println("interface examples")
	c := circle{radius: 5}
	r := rect{width: 10, height: 20}
	shapeStarts(&c)
	shapeStarts(&r)
}

// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
