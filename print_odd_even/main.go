package main

import (
	"fmt"
)

/*
Bài 3: In số chẵn, số lẻ xen kẽ
Viết một chương trình dùng 2 goroutines:

Một goroutine in ra số chẵn từ 0 đến 20.
Một goroutine in ra số lẻ từ 1 đến 19.
Yêu cầu:

Sử dụng channel để đảm bảo luồng in xen kẽ nhau.
*/

func main() {
	evenChan := make(chan bool)
	oddChan := make(chan bool)

	// Goroutine in số chẵn
	go func() {
		for i := 0; i <= 20; i += 2 {
			<-evenChan // Chờ tín hiệu từ channel evenChan
			fmt.Println("Even:", i)
			oddChan <- true // Gửi tín hiệu cho oddChan
		}
	}()

	// Goroutine in số lẻ
	go func() {
		for i := 1; i < 20; i += 2 {
			<-oddChan // Chờ tín hiệu từ channel oddChan
			fmt.Println("Odd:", i)
			evenChan <- true // Gửi tín hiệu cho evenChan
		}
	}()

	// Bắt đầu in bằng cách kích hoạt evenChan
	evenChan <- true

	// Đợi người dùng nhấn phím để kết thúc chương trình
	fmt.Scanln()
}
