package main

import (
	"fmt"
	"strings"
)

func main() {
	// result := strings.Contains("BabelCoder", "bel")
	// result := strings.Count("สวัสดี โชคดีนะ", "ดี")
	// result := strings.HasPrefix("Hello World", "Hell")
	// result := strings.HasSuffix("Hello World", "ld")
	// result := strings.Join([]string{"สวัสดี", "ชาวโลก"}, "-")
	result := strings.ToUpper("hello")

	fmt.Println("result ::", result)
}
