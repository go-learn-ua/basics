package main

import "fmt"

func main() {
	amIInTheBackeryShop := true
	isShopOpen := true
	isBapExists := true
	//doYouHaveYourOwnChief := true

	if amIInTheBackeryShop && isShopOpen && isBapExists {
		fmt.Println("Я можу купити булочку")
		return
	}

	fmt.Println("Я не можу купити булочку")
}
