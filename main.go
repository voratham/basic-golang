package main

import "fmt"

type enrollment struct {
	semaster string
	courses  []string
}

// method receiver as value receiver
func (e enrollment) courseAt(index uint) string {
	return e.courses[index]
}

// method receiver as pointer receiver
func (e *enrollment) addCourse(course string) {
	e.courses = append(e.courses, course)

}

func main() {
	e := enrollment{semaster: "1/63", courses: []string{"Java", "C#"}}

	// e.courses = append(e.courses, "C++", "Javascript")

	// equal (&e).addCourse("Golang") because golang complier smart and auto convert
	(e).addCourse("Golang")

	result := e.courseAt(0)
	fmt.Println("index ::", result)

	fmt.Printf("%+v", e)

}
