package main

import "fmt"

func main() {
	exerciseNames := []string{"Підтягування", "Бурпі", "Присідання"}
	for exerciseIndex := range exerciseNames {
		fmt.Println("Прийшов час вправи", exerciseNames[exerciseIndex])
		for count := 0; count < 10; count++ {
			fmt.Println(exerciseNames[exerciseIndex], count+1)
		}
	}
}
