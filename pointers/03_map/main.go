package main

import "fmt"

func main() {
	m := make(map[string]string)
	updateMap(m)
	fmt.Println(m)
}

func updateMap(m map[string]string) {
	m["1"] = "1"
	m["2"] = "2"
	m["3"] = "3"
	m["4"] = "4"
	m["5"] = "5"
	m["6"] = "6"
	m["7"] = "7"
	m["8"] = "8"
	m["9"] = "9"
	m["10"] = "10"
}
