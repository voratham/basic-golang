package main

import "fmt"

func main() {
	// const (
	// 	red   = 0
	// 	green = 1
	// 	blue  = 2
	// )

	// const (
	// 	red = iota + 1
	// 	green
	// 	blue
	// )

	// fmt.Println(red, green, blue)

	const (
		sun = iota + 1
		mon
		tue
		_
		_
		_
		sat
	)
	fmt.Println(sun, mon, tue, sat)
}
