package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Println("Hello World!")
	fmt.Println(plus(a, b))

}

func plus(a int, b int) int {
	a = 500
	b = 500

	return a + b

}
