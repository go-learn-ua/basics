package main

import "fmt"

func main() {
	myText := "мій текст більший"
	myFriendsText := "текст"

	isMyTextHasMoreChars := myText > myFriendsText
	fmt.Println("Чи мій текст більший? ", isMyTextHasMoreChars)
}
