package main

import "fmt"

func main() {
	exerciseNames := []string{"Підтягування", "Бурпі", "Присідання"}
	for exerciseIndex := 0; exerciseIndex < len(exerciseNames); exerciseIndex++ {
		fmt.Println("Прийшов час вправи", exerciseNames[exerciseIndex])
		for count := 0; count < 10; count++ {
			fmt.Println(exerciseNames[exerciseIndex], count+1)
		}
	}
}
