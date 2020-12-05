package main

import "fmt"

func main() {
	// []byte, byte = unit8 (0 - 255)
	str := "Hello World"
	firstLetter := str[0]
	fmt.Println(firstLetter) // ASCII => H

	// we cannot set variable str[0] = "A" because It is byte

	fmt.Println("after convert : ", string(firstLetter))
	fmt.Println("len", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

}
