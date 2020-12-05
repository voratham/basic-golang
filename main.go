package main

import "fmt"

type generator interface {
	generate()
}

type pdf struct {
	content string
}

func (p pdf) generate() {
	fmt.Println("Generating...")
}

func main() {
	var gen generator
	gen = pdf{content: "My PDF"}
	gen.generate()
}
