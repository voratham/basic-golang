package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// rune => int32 = unicode
	str := "สวัสดีชาวโลก"

	// ch := "ส"
	// fmt.Println(([]byte(ch)))

	word := []rune(str)
	fmt.Println("rune", word[0], "string()", string(word[0]), "len", len(str))
	fmt.Println("utf8 rune count in string :", utf8.RuneCountInString(str))

	// we should use for range because golang will loop autu converted
	for _, char := range str {
		fmt.Print(string(char))
	}

	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(string(str[i]))
	// }

}
