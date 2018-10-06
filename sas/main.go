package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter number of values, then the values:")
	fmt.Println("n")
	fmt.Println("n1 n2 n3 . . .")
	fmt.Println("To end, enter RETURN")
	size := readSize()
	for size > 0 {
		arr := make([]int, size)
		readElements(arr)
		sum := SimpleArraySum(arr)
		fmt.Println("sum:", sum)
		size = readSize()
	}
}

// SimpleArraySum returns the sum of slice of integers ints
func SimpleArraySum(ints []int) (sum int) {
	sum = 0
	for _, v := range ints {
		sum += v
	}
	return
}

func readSize() int {
	var n int
	fmt.Scanln(&n)
	return n
}
func readElements(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
}
