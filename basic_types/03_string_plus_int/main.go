package main

import "fmt"

func main() {
	var Liubomyr, Yura, Bohdan, Oleh int

	Liubomyr = 4
	Yura = 3
	Bohdan = 4
	Oleh = 25

	hinkali := Liubomyr + Yura + Bohdan + Oleh
	order := fmt.Sprintf("Замовлення = %s, кількість = %d", "хінкалі", hinkali)

	fmt.Println(order)
}
