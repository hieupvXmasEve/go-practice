package main

import (
	"fmt"
	"sync"
)

/*
Bài 5: Truy cập dữ liệu chia sẻ
Viết chương trình có một biến đếm (counter) được chia sẻ.

Có 3 goroutines: mỗi goroutine tăng giá trị biến counter 1000 lần.
Yêu cầu:

Sử dụng sync.Mutex hoặc sync/atomic để tránh race condition.
*/

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex // Tạo Mutex

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()   // Khoá Mutex trước khi thay đổi giá trị của counter
				counter++   // Thay đổi giá trị của counter
				mu.Unlock() // Mở khoá Mutex sau khi thay đổi
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
