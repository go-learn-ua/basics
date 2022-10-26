package main

import "fmt"

var menu = []string{"хінкалі", "хачапурі", "хaрчо", "чібупелі", "чахохбілі", "чурчхела", "шаурма"}

func main() {
	result, err := order("чібупелі", 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = order("чорний трюфель маринований у віні шевал бланк", 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func order(dish string, amount uint) (string, error) {
	for _, item := range menu {
		if dish == item {
			return fmt.Sprintf("Замовлення = %s, кількість = %d", dish, amount), nil
		}
	}

	return "", fmt.Errorf("блюдо не знайдено")
}
