package shaurma

import (
	"fmt"
	"recipes/cabbage"
	"recipes/meat"
)

var Name = "Шаурма"

type Shaurma struct {
	Name         string
	FreshCabbage cabbage.Cabbage
	Meat         meat.Meat
}

func EatUp(s Shaurma) {
	fmt.Println(fmt.Sprintf("Смачного, %s  зайшла як додому", s.Name))
	fmt.Println(fmt.Sprintf("Складник: %s  вагою: %d", s.FreshCabbage.Name, s.FreshCabbage.Amount))
	fmt.Println(fmt.Sprintf("Складник: %s  вагою: %d", s.Meat.Name, s.Meat.Amount))
}
