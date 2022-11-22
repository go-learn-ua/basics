package main

import "fmt"

func main() {
	cookiesPrice := 20
	myMoney := 20

	canByCookies := myMoney >= cookiesPrice
	fmt.Println("Чи можу я купити печення: ", canByCookies)

	canByCookies = cookiesPrice <= myMoney
	fmt.Println("Чи можу я купити печення: ", canByCookies)
}
