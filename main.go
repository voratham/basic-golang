package main

import (
	"fmt"
	"runtime"
)

func main() {
	// num := 20
	// if num%2 == 0 {
	// 	fmt.Println("Even")
	// } else {
	// 	fmt.Println("Odd")
	// }

	if os := runtime.GOOS; os == "darwin" {
		fmt.Println("You is macos 🖥")
	} else if os == "windows" {
		fmt.Println("You is window 😎")
	} else {
		fmt.Println("You is linux 👨‍💻 ?")
	}
}
