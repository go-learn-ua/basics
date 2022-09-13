package main

import "fmt"

func main() {
	exerciseNames := []string{"Підтягування", "Бурпі", "Присідання"}
	for _, exerciseName := range exerciseNames {
		fmt.Println("Прийшов час вправи", exerciseName)
		for count := 0; count < 10; count++ {
			fmt.Println(exerciseName, count+1)
		}
	}
}
