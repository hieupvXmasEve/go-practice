package main

import (
	"fmt"
	"os"
)

/*
Bài tập 4: Hệ thống log đa dạng
Tạo một interface Logger với phương thức Log(message string).
Tạo các struct triển khai interface Logger:
ConsoleLogger (in ra màn hình).
FileLogger (ghi ra file).
Viết một hàm sử dụng interface Logger để ghi log từ một slice message vào cả console và file.

*/

type Logger interface {
	Log(message string)
}

// Struct ConsoleLogger để ghi log ra console
type ConsoleLogger struct{}

// Triển khai phương thức Log cho ConsoleLogger
func (c ConsoleLogger) Log(message string) {
	fmt.Printf("Console log: %s\n", message)
}

// Struct FileLogger để ghi log ra file
type FileLogger struct {
	FileName string
}

// Triển khai phương thức Log cho FileLogger
func (f FileLogger) Log(message string) {
	// Mở file để ghi (hoặc tạo file nếu chưa có)
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Lỗi khi mở file: %v\n", err)
		return
	}
	defer file.Close()

	// Ghi log vào file
	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Printf("Lỗi khi ghi log vào file: %v\n", err)
		return
	}

	fmt.Printf("File log: %s (đã ghi vào %s)\n", message, f.FileName)
}

// Hàm xử lý danh sách log với nhiều loại logger
func ProcessLogs(loggers []Logger, messages []string) {
	for _, message := range messages {
		for _, logger := range loggers {
			logger.Log(message)
		}
	}
}

// Hàm main kiểm tra chương trình
func main() {
	// Tạo loggers
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{FileName: "log.txt"}

	// Danh sách thông báo log
	messages := []string{
		"Hệ thống khởi động thành công.",
		"Người dùng đăng nhập.",
		"Lỗi: Không thể kết nối cơ sở dữ liệu.",
	}

	// Ghi log bằng cả hai loại logger
	loggers := []Logger{consoleLogger, fileLogger}
	ProcessLogs(loggers, messages)

	fmt.Println("Hoàn tất ghi log.")
}
