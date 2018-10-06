package main

import (
	"fmt"
)

func main() {
	var arr []int
	sum := SimpleArraySum(arr)
	fmt.Println(sum)
}

// SimpleArraySum returns the sum of slice of integers ints
func SimpleArraySum(ints []int) (sum int) {
	sum = 0
	for _, v := range ints {
		sum += v
	}
	return
}
