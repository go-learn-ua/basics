package main

import "fmt"

func main() {
	number := 7
	switch number {
	case 1:
		fmt.Println("Обрані львівські цляцки")
	case 2:
		fmt.Println("Сьогоді буде котлета по-київськи")
	case 3, 4:
		fmt.Println("Вітаю, потавськи галушки майже готові!")
	default:
		fmt.Println("Одеську подвійну асорті шаурму, будь ласка")
	}
}
