package main

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		min := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > min {
			arr[j+1] = arr[j]
			arr[j] = min
			j--
		}
	}
	return arr
}
