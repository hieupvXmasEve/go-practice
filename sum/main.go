package main

import (
	"fmt"
	"sync"
)

/*
Bài 2: Tính tổng trong mảng lớn
Viết chương trình chia một mảng số nguyên lớn thành 2 phần và sử dụng 2 goroutines để tính tổng mỗi phần.

Yêu cầu:

Kết quả cuối cùng là tổng của cả mảng.
Dùng sync.WaitGroup hoặc channel để gom kết quả.
*/

func sum(list []int, s chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, i := range list {
		sum += i
	}
	s <- sum
}
func main() {
	var wg sync.WaitGroup

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := make(chan int)
	wg.Add(2)
	go sum(list[:len(list)/2], s, &wg)
	go sum(list[len(list)/2:], s, &wg)
	go func() {
		wg.Wait() // Đợi cả 2 goroutine hoàn thành
		close(s)  // Đóng channel sau khi hoàn tất
	}()
	total := 0
	for partialSum := range s {
		total += partialSum
	}
	fmt.Println("Tổng của mảng:", total)

}
