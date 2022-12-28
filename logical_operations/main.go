package main

import "fmt"

func main() {
	amIInTheBackeryShop := true
	isShopOpen := false
	isBapExists := true
	doYouHaveYourOwnChief := true

	if (amIInTheBackeryShop && isShopOpen && isBapExists) || doYouHaveYourOwnChief {
		fmt.Println("Я можу купити булочку")
		return
	}

	fmt.Println("Я не можу купити булочку")
}
