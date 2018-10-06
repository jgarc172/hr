package main

import (
	"fmt"
)

func main() {
	var a, b, sum int
	fmt.Println("Enter two integer values. To stop click ENTER:")
	_, err := fmt.Scanln(&a, &b)
	for err == nil {
		sum = Add(a, b)
		fmt.Println("sum:", sum)
		_, err = fmt.Scanln(&a, &b)
	}
}

// Add is equivalent to the + operator
func Add(a, b int) (sum int) {
	sum = a + b
	return
}
