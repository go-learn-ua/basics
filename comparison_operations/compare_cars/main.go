package main

import "fmt"

func main() {
	expectedTaxi := "Ферарі 212"
	carNearToMe := "Ферарі 212"

	isItTaxi := expectedTaxi == carNearToMe
	fmt.Println("чи то моє таксі", isItTaxi)

	isItNotTaxi := expectedTaxi != carNearToMe
	fmt.Println("чи то НЕ моє таксі? ", isItNotTaxi)
}
