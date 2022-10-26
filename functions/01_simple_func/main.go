package main

import "fmt"

func main() {
	hi("Olenka")
	hi("Stepan")
	hi("Taras")
}

func hi(name string) {
	fmt.Printf("Hi %s!", name)
}
