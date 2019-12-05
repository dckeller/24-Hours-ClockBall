package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perim() float32
}

type rectangle struct {
	width  float32
	height float32
}

type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	return r.width * r.height
}

func (r rectangle) perim() float32 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}
	measure(r)
	measure(c)
}
