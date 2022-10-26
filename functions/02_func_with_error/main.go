package main

import "fmt"

func main() {
	err := hi("Olenka")
	if err != nil {
		fmt.Println(err)
	}
	err = hi("Stepan")
	if err != nil {
		fmt.Println(err)
	}
	err = hi("Taras")
	if err != nil {
		fmt.Println(err)
	}
}

func hi(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("вказано пусте ім'я")
	}

	fmt.Printf("Hi %s!", name)
	return nil
}
