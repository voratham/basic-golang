package main

import "fmt"

// 0x1111 20
func main() {
	// i := 20
	// fmt.Println(i, &i)
	i := 20
	var p *int

	p = &i

	fmt.Println("[1] point ::", p, "print point real val :", *p)

	i = 30

	fmt.Println("[2] point ::", p, "print point real val :", *p)
}
