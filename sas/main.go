package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter number of values, then the values:")
	fmt.Println("n")
	fmt.Println("n1 n2 n3 . . .")
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sum := SimpleArraySum(arr)
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
