package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s shape
	s = circle{radius: 10}

	switch s.(type) {
	case circle:
		fmt.Println("circle")
	case rectangle:
		fmt.Println("rectangle")
	}

	// s,ok = circle{radius: 10}
	// fmt.Println("ok :", ok)
	// fmt.Printf("%+v", c)
}
