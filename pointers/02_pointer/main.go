package main

import "fmt"

func main() {
	value := 1

	ptr1 := &value
	fmt.Println(ptr1)

	ptr2 := &value
	fmt.Println(ptr2)

	*ptr2 = 2
	fmt.Println(value, *ptr1, *ptr2)

	value2 := 2
	updateInt(&value2)
	fmt.Println(value2)

}

func updateInt(ptr *int) {
	*ptr = 23
}
