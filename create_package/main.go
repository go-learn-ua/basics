package main

import (
	"recipes/cabbage"
	"recipes/meat"
	"recipes/shaurma"
)

func main() {
	first := shaurma.Shaurma{
		Name: "Київська",
		Meat: meat.Meat{
			Name:   "курка",
			Amount: 500,
		},
		FreshCabbage: cabbage.Cabbage{
			Name:   "білокачанна",
			Amount: 100,
		},
	}

	shaurma.EatUp(first)
}
