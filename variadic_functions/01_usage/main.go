package main

import "fmt"

func main() {
	fmt.Println(1, 2, "string", 3.2)

	s := fmt.Sprintf("основних уроків %s буде %d", "go", 20)

	fmt.Println(s)
}
