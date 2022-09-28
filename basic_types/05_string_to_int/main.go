package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("50 чогось там")
	fmt.Println(err)
	fmt.Println(i)

	i, err = strconv.Atoi("50")
	fmt.Println(err)
	fmt.Println(i)
}
