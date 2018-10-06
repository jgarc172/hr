package main

import (
	"fmt"
)

func main() {
	var sum int
	a := 4
	b := 3
	sum = solveMeFirst(a, b)
	fmt.Println(sum)
}

func solveMeFirst(a, b int) (sum int) {
	sum = a + b
	return
}
