package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 0, 5)

	fmt.Println("s1 :", s1, "len(s1) :", len(s1), " cap(s1):", cap(s1))
	fmt.Println("s2 :", s2, "len(s2) :", len(s2), " cap(s2):", cap(s2))

	s2 = append(s2, 1, 2, 3, 4, 5)
	// create s3
	s3 := append(s2, 6)
	fmt.Println("[2] -s2 :", s2, "len(s2) :", len(s2), " cap(s2):", cap(s2))
	fmt.Println("s3 :", s3, "len(s3) :", len(s3), " cap(s3):", cap(s3))

	// change value from s3 poc mutate
	s3[0] = 10

	fmt.Println("[3] -s2 :", s2, "len(s2) :", len(s2), " cap(s2):", cap(s2))
	fmt.Println("[2] -s3 :", s3, "len(s3) :", len(s3), " cap(s3):", cap(s3))

}
