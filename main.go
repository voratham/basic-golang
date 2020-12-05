package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, ok := strconv.ParseFloat("3.14", 64)
	fmt.Println("a", a, "ok:", ok)

	e, ok := strconv.ParseInt("0110", 2, 64) // second parameter  is set number unit
	fmt.Println("e", e, "ok:", ok)

	e2, ok := strconv.ParseInt("120", 10, 64)
	fmt.Println("e2", e2, "ok:", ok)

	i, ok := strconv.ParseUint("124", 10, 64)
	fmt.Println("i", i, "ok:", ok)

	o, ok := strconv.Atoi("65")
	fmt.Println("o", o, "ok:", ok)

	w := strconv.Itoa(65)

	fmt.Println("w", w, "ok:", ok)

}
