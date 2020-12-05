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
	author  user
}

func main() {
	// u := user{"Somechai", 21}
	// u := user{name: "John", age: 21}
	// u.age = 22

	a := article{
		// title:   "test ",
		// excerpt: "Excerpt",
		author: user{name: "John", age: 21},
	}

	a.title = "The Title"

	fmt.Printf("%+v", a)
}
