package sort

// tìm kiếm phần tử nhỏ nhất trong mảng sau đó mới đổi vị trí. (Thường nhầm lẫn ở việc là hoán đổi luôn trong vòng lặp)
func SelectionSort(arr []int) []int {
	n := len(arr)
	var i, j, min int

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
