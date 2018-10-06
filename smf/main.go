package main

import (
	"fmt"
)

func main() {
	var a, b, sum int
	fmt.Scanf("%v\n%v", &a, &b)
	sum = solveMeFirst(a, b)
	fmt.Println(sum)
}

func solveMeFirst(a, b int) (sum int) {
	sum = a + b
	return
}
