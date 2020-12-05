package main

import "fmt"

type user struct {
	name string
	age  uint
}

type article struct {
	title   string
	excerpt string
	body    string
	user
}

func main() {

	a := article{
		title:   "test ",
		excerpt: "Excerpt",
		body:    "string",
		user:    user{name: "John", age: 21},
	}

	a.title = "The Title"

	fmt.Printf("%+v", a)
	fmt.Printf("%+v", a.age) // promote filed can call field user
}
