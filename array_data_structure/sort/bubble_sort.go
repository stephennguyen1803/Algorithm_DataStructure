package sort

// Sắp xếp nổi bọt có điểm đặt trưng là hoán đổi 2 thành phần kế tiếp nhau và đưa về cố định ở một điểm thường
// là điểm cuối của mảng. Điểm đặt biệt dễ sai đó là j khởi động từ 0 (dễ mắc lỗi là gắn bằng i) (vì cần so sánh 0 và 1 trước). Đồng thời
// điều kiện chấm dứt vòng lặp j phải kiểm tra j < n - 1 -i vì j bắt đầu từ i
// thêm biến check để tối ưu buble sort => nếu cả quá trình so sánh j và j+1 mà không có hoán đổi vị trí => Tức là mảng đã dc
// sắp xếp
func BubbleSort(arr []int) []int {
	n := len(arr)
	var i, j, t int

	for i = 0; i < n-1; i++ {
		check := false
		for j = 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				t = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
				check = true
			}
		}

		if !check {
			return arr
		}
	}

	return arr
}
