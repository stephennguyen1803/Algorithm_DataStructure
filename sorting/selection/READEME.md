#1.Selection Sort
## Ý Chính Bao Gồm 2 mảng con dựa trên mảng chính cần sắp xếp
1) Ý tưởng: Chọn phần tử nhỏ nhất và đưa nó về đầu mảng. Lặp lại quá trình với phần còn lại của mảng.
1) Mảng con để chứa các phần tử được sắp xếp.
2) Mảng con còn lại chưa được sắp xếp
## flowchart: https://media.geeksforgeeks.org/wp-content/cdn-uploads/20220203094305/Selection-Sort-Flowhchart.png

### Độ phức tạp thuật toán:
O(n^2)
1. Sử dụng vòng lặp từ 1 -> đến n - 1 (vì ta sẽ cần sử dụng giá trị i + 1) : sử dụng phần tử i để gắn giá trị cần so sánh (độ phức tạp O(n))
2. Chạy vòng thứ 2 từ j = i +1 đến n: so sánh hai giá trị a[i] và a[j]. Nếu a[i] > a[j] thì đổi chổ hai phần tử (độ phức tạp O(n))
3. vì 2 vòng lặp lồng nhau: => độ phức tạp của thuật toán sẽ là B(O) = O(n) * O(n) = O(n^2)
