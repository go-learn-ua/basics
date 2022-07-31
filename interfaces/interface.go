package main

import "fmt"

type boiler interface {
	boil()
}

type kettle struct{}

func (kettle) boil() {
	fmt.Println("чайник: вода закіпячена")
}

type pot struct{}

func (pot) boil() {
	fmt.Println("баняк: вода закіпячена")
}

func makeMivina(b boiler) {
	fmt.Println("взяти тарілку")
	fmt.Println("пололжити мівіну в тарілку")

	b.boil()

	fmt.Println("залити мівіну водою")
	fmt.Println("накрити кришкою")
}

func main() {
	makeMivina(pot{})
}
