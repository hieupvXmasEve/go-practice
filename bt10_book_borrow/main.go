package main

import (
	"errors"
	"fmt"
)

/*
	Bài 5: Quản lý thư viện sách

	1.	Tạo interface Book với các phương thức:
	•	GetTitle() string: Lấy tiêu đề sách.
	•	Borrow() error: Mượn sách (kiểm tra sách có sẵn không).
	•	Return(): Trả sách.
	2.	Tạo hai loại sách:
	•	PhysicalBook với thuộc tính Title, Author, AvailableCopies.
	•	EBook với thuộc tính Title, Author, LicenseKey.
	3.	Viết chương trình quản lý danh sách sách trong thư viện:
	•	Cho phép mượn và trả sách.
	•	Hiển thị trạng thái của từng cuốn sách (số bản còn lại hoặc trạng thái đã được mượn).
*/

// Định nghĩa interface Book
type Book interface {
	GetTitle() string
	Borrow() error
	Return()
	GetStatus() string
}

// Định nghĩa PhysicalBook
type PhysicalBook struct {
	Title           string
	Author          string
	AvailableCopies int
}

func (p *PhysicalBook) GetTitle() string {
	return p.Title
}

func (p *PhysicalBook) Borrow() error {
	if p.AvailableCopies > 0 {
		p.AvailableCopies--
		return nil
	}
	return errors.New("không còn bản sách vật lý nào để mượn")
}

func (p *PhysicalBook) Return() {
	p.AvailableCopies++
}

func (p *PhysicalBook) GetStatus() string {
	return fmt.Sprintf("Sách vật lý: %d bản có sẵn", p.AvailableCopies)
}

// Định nghĩa EBook
type EBook struct {
	Title      string
	Author     string
	LicenseKey string
	IsBorrowed bool
}

func (e *EBook) GetTitle() string {
	return e.Title
}

func (e *EBook) Borrow() error {
	if !e.IsBorrowed {
		e.IsBorrowed = true
		return nil
	}
	return errors.New("sách điện tử đã được mượn")
}

func (e *EBook) Return() {
	e.IsBorrowed = false
}

func (e *EBook) GetStatus() string {
	if e.IsBorrowed {
		return "Sách điện tử: đã được mượn"
	}
	return "Sách điện tử: có sẵn"
}

// Hàm chính quản lý thư viện
func main() {
	// Tạo danh sách sách
	library := []Book{
		&PhysicalBook{Title: "Lập trình Go", Author: "John Doe", AvailableCopies: 3},
		&PhysicalBook{Title: "Cơ sở dữ liệu", Author: "Jane Smith", AvailableCopies: 1},
		&EBook{Title: "Học thuật toán", Author: "Alice Johnson", LicenseKey: "ABC123", IsBorrowed: true},
	}

	// Hiển thị trạng thái ban đầu của sách
	fmt.Println("Trạng thái sách ban đầu:")
	for _, book := range library {
		fmt.Println(book.GetTitle(), "-", book.GetStatus())
	}

	// Mượn sách
	fmt.Println("\nMượn sách:")
	for _, book := range library {
		err := book.Borrow()
		if err != nil {
			fmt.Printf("Không thể mượn \"%s\": %v\n", book.GetTitle(), err)
		} else {
			fmt.Printf("Đã mượn \"%s\" thành công.\n", book.GetTitle())
		}
	}

	// Hiển thị trạng thái sau khi mượn
	fmt.Println("\nTrạng thái sách sau khi mượn:")
	for _, book := range library {
		fmt.Println(book.GetTitle(), "-", book.GetStatus())
	}

	// Trả sách
	fmt.Println("\nTrả sách:")
	for _, book := range library {
		book.Return()
		fmt.Printf("Đã trả \"%s\" thành công.\n", book.GetTitle())
	}

	// Hiển thị trạng thái sau khi trả
	fmt.Println("\nTrạng thái sách sau khi trả:")
	for _, book := range library {
		fmt.Println(book.GetTitle(), "-", book.GetStatus())
	}
}
