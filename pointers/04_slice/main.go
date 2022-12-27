package main

import "fmt"

func main() {
	s := make([]int, 2)
	s[0] = 0
	s[1] = 0

	updateSlice(s)
	fmt.Println(s)

	appendToSlice(&s)
	fmt.Println(s)
}

func updateSlice(s []int) {
	s[0] = 1
	s[1] = 1
}

func appendToSlice(s *[]int) {
	*s = append(*s, 1)
	*s = append(*s, 1)
}
