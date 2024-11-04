package main

import (
	sortdata "array-structure/sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Nhập số phần tử của mảng
	fmt.Print("Nhập số phần tử của mảng: ")
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(input[:len(input)-1])

	// Khởi tạo mảng
	arr := make([]int, n)

	// Nhập các phần tử của mảng
	for i := 0; i < n; i++ {
		fmt.Printf("Nhập phần tử thứ %d: ", i+1)
		input, _ := reader.ReadString('\n')
		arr[i], _ = strconv.Atoi(input[:len(input)-1])
	}

	// In ra mảng đã nhập
	fmt.Println("Mảng đã nhập:", arr)

	//sắp xếp mảng lựa chọn phần tử rồi hoán đổi - thường phần tử đầu mảng sẽ cố định
	sortArr := sortdata.SelectionSort(arr)
	//sắp xếp mảng sử dụng việc hoán đổi hai phần từ liền kề, phần tử ở cuối mảng sẽ cố định
	sortBubbleArr := sortdata.BubbleSort(arr)

	fmt.Println("Mảng sau khi sắp xếp lựa chọn: ", sortArr)
	fmt.Println("Mảng sau khi sắp xếp bong bóng: ", sortBubbleArr)
}
