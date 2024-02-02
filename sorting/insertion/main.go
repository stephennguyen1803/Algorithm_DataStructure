package main

import "fmt"

func main() {
	arr := []int{3, 2, 1, 5, 4}
	sortArr := InsertionSort(arr)

	fmt.Println(sortArr)
}
