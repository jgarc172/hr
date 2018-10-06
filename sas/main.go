package main

import (
	"fmt"
)

func main() {
	var arr []int
	fmt.Println("length of arr:", len(arr))
	sum := SimpleArraySum(arr)
	fmt.Println("sum:", sum)
	arr = []int{1, 2, 3}
	fmt.Println("length of arr:", len(arr))
	sum = SimpleArraySum(arr)
	fmt.Println("sum:", sum)
}

// SimpleArraySum returns the sum of slice of integers ints
func SimpleArraySum(ints []int) (sum int) {
	sum = 0
	for _, v := range ints {
		sum += v
	}
	return
}
