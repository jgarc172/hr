package main

import (
	"fmt"
)

func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)
	sum = Add(a, b)
	fmt.Println(sum)
}

// Add is equivalent to the + operator
func Add(a, b int) (sum int) {
	sum = a + b
	return
}
