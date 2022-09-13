package main

import "fmt"

func main() {
	count := 0

	for {
		if count < 10 {
			fmt.Println("Присідання", count+1)
		} else {
			break
		}
		count++ // count = count + 1
	}
}
