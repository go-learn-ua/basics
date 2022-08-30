package main

import "fmt"

func main() {
	var t = -15

	if t > -20 {
		fmt.Println("рятуймося, глобальне потепління!")
	} else if t == -20 {
		fmt.Println("все норм, спостерігаємо, п'ємо смузі")
	} else {
		fmt.Println("купуй пічку!")
	}
}
