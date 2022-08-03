package main

import "fmt"

func main() {
	var emptyAge int
	fmt.Println("Пуста змінна", emptyAge)

	age := 1
	fmt.Println("Початковий вік автора", age)

	age = 25
	fmt.Println("Вік автора", age)

	age = 26
	fmt.Println("Змінений вік автора", age)
}
