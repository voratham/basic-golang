package main

import (
	"course-go/shape"
	"fmt"
)

func main() {

	circle := shape.Circle{Radius: 10}
	fmt.Printf("%+v", circle.Area())

}
