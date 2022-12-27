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
}
