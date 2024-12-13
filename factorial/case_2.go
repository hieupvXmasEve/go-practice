package main

import (
	"fmt"
)

/*
Bài 4: Tính giai thừa song song
Viết một chương trình tính giai thừa của các số từ 1 đến 5.

Mỗi số được tính bởi một goroutine khác nhau.
Yêu cầu:

Kết quả giai thừa của mỗi số được in ra sau khi tính xong.
Dùng channel để thu kết quả từ các goroutines.
*/

// Hàm tính giai thừa
func factorial2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Tạo một channel để thu kết quả
	results := make(chan string)

	// Chạy goroutines để tính giai thừa cho các số từ 1 đến 5
	for i := 1; i <= 5; i++ {
		go func(num int) {
			// Gửi kết quả vào channel
			results <- fmt.Sprintf("Giai thừa của %d là %d", num, factorial2(num))
		}(i)
	}

	// Nhận kết quả từ channel và in ra
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}

}
