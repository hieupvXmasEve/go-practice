package main

import "sync"

/*
Bài 1: In ra chuỗi song song
Viết một chương trình sử dụng 2 goroutines để in ra 2 chuỗi khác nhau ("A" và "B") 10 lần mỗi chuỗi.

Yêu cầu:

- Sử dụng go để chạy goroutine.
- Dùng sync.WaitGroup để đảm bảo tất cả goroutines hoàn thành.
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			println("A")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			println("B")
		}
		wg.Done()
	}()

	wg.Wait()
}
