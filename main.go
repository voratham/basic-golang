package main

import "fmt"

func f() {
	panic("from f")
}

func g() {
	f()
}

func h() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "Recover")

		}
	}()

	g()
}

func main() {
	h()
}
