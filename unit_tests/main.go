package main

import (
	"fmt"
)

var menu = []string{"хінкалі", "хачапурі", "шаурма"}

//create function with amount alements in menu

func main() {
	result, err := Order("чібупелі", 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = Order("чорний трюфель маринований у віні шевал бланк", 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// create function with amount alements in menu
func Order(dish string, amount uint) (string, error) {
	for _, item := range menu {
		if dish == item {
			return fmt.Sprintf("Замовлення = %s, кількість = %d", dish, amount), nil
		} else {
			continue
		}
	}

	return "", fmt.Errorf("блюдо не знайдено")
}
