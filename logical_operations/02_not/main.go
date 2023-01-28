package main

import "fmt"

func main() {
	amIInTheBackeryShop := true
	isShopOpen := true
	isBapExists := true
	doYouHaveYourOwnChief := false

	if amIInTheBackeryShop && isShopOpen && isBapExists && !doYouHaveYourOwnChief {
		fmt.Println("Я можу купити булочку")
		return
	}

	fmt.Println("Я не можу купити булочку")
}
