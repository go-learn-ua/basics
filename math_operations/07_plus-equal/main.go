package main

import "fmt"

func main() {
	bradley := 40
	marder := 30
	bradley += marder
	fmt.Println(bradley)

	bradley = 40
	marder = 30
	bradley = bradley + marder
	fmt.Println(bradley)
}
