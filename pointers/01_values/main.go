package main

import "fmt"

func main() {
	value1 := 1
	value2 := value1
	value2 += 4
	fmt.Println("Змінні", "Значення 1", value1, "Значення 2", value2)

	value3 := 3
	updateNumber(value3)
	fmt.Println(value3)
}

func updateNumber(value3 int) {
	value3 = 2
}
