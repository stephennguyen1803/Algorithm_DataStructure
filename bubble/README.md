### Bubble Sort
##1. Thuật toán dựa trên so sánh hai cặp số liền kề:
##2. Chú ý quan trọng, các phần tử được sắp xếp sẽ được xếp lên cuối mảng (vòng lặp sau sẽ không tính đến phần tử cuối cùng này), cứ xong một vòng lặp thì phần tử được sắp xếp lại tăng lên 1 (cuối mảng)
```
[5 2 6 8 1 3]
#lần 1:
5 > 2 => đổi chổ 5 và 2 => [2 5 6 8 1 3]
5 < 6 => ko đổi => [2 5 6 8 1 3]
6 < 8 => ko đổi => [2 5 6 8 1 3]
8 > 1 => đổi chổ 8 và 1 => [2 5 6 1 8 3]
8 > 3 => đổi chổ 8 và 3 => [2 5 6 1 3 8] (phần tử 8 đã được sắp xếp nên lần sau sẽ ko check nữa)
#lần 2:
2 < 5 => ko đổi => [2 5 6 1 3 8]
5 < 6 => ko đổi => [2 5 6 1 3 8]
6 > 1 => đổi chổ 6 và 1 => [2 5 1 6 3 8]
6 > 3 => đổi chổ 6 và 3 => [2 5 1 3 6 8] (ngừng tại đây vì 8 đã đc sắp xếp từ lần 1)
#lần 3:
2 < 5 => ko đổi => [2 5 1 3 6 8]
5 > 1 => đổi chổ 5 và 1 => [2 1 5 3 6 8]
5 > 3 => đổi chổ 5 và 3 => [2 1 3 5 6 8] (ngừng tại đây vì 6, 8 đã dc sắp xếp từ lần 1 và 2)
#lần 4:
2 > 1 => đổi chổ 2 và 1 => [1 2 3 5 6 8] 
2 < 3 => ko đổi chổ => [1 2 3 5 6 8] (ngừng tại đây vì 5, 6, 8 đã đc sắp xếp từ lần 1, 2 và lần 3)
#lần 5:
1 < 2 => ko đổi => [1 2 3 5 6 8] (ngừng tại đây vì 3, 5, 6, 8 đã đc sắp xếp từ lần 1, 2, 3 và lần 4)
==> In ra kết quả mảng đã được sắp xếp.
```
##3. Độ phức tạp thuật toán là O(n^2)
##4. Thuật toán bubble sort optimize dựa trên ý tưởng nếu vòng lặp kiểm tra thứ n ko có đổi phần tử nào thì ko cần chạy thêm vòng lặp n+1 để kiểm tra nữa mà ngừng luôn để tiết kiệm bộ nhớ và giải thuật.
```
[1 8 3 5 6]
# khởi tạo flag check = true) để chạy kiểm tra
#lần 1 (nếu flag check = true mới chạy)
set flag check = false (trong trường hợp mảng ko có đổi chỗ thì => mảng đã đc sắp xếp nên ko cần tiếp tục check)
1 < 8 => ko đổi chổ => [1 8 3 5 6]
8 > 3 => đổi chổ 8 và 3 => [1 3 8 5 6] (có sự thay đổi giá trị nên flag check = true)
8 > 5 => đổi chổ 8 và 5 => [1 3 5 8 6]
8 > 6 => đổi chổ 8 và 6 => [1 3 5 6 8] (phần tử 8 đã đc sắp xếp nên lần sau sẽ ko check nữa)
#lần 2 (nếu flag check = true mới chạy)
set flag check = false (trong trường hợp mảng ko có đổi chỗ thì => mảng đã đc sắp xếp nên ko cần tiếp tục check)
1 < 3 => ko đổi => [1 3 5 6 8]
3 < 5 => ko đổi => [1 3 5 6 8]
5 < 6 => ko đổi => [1 3 5 6 8] (ngừng tại đây vì phần tử 8 đã được sắp xếp nên lần sau sẽ ko check nữa)
#lần 3 flag check = false ngừng vòng lặp in kêt quả (do mảng đã đc sắp xếp ko có thêm phần tử nào cần hoán đổi)
```
