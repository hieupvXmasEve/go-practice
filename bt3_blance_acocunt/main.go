package main

import (
	"fmt"
)

/*
Bài tập 3: Quản lý tài khoản ngân hàng
Tạo struct Account với các thuộc tính:
ID, Name, Balance
Viết các phương thức sử dụng pointer để:
Nạp tiền vào tài khoản.
Rút tiền (nếu số dư đủ).
Hiển thị thông tin tài khoản.
Tạo thêm một interface Transaction với phương thức Process() và triển khai cho các loại giao dịch (nạp tiền, rút tiền).
*/

type Account struct {
	ID      int
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if a.Balance < amount {
		fmt.Println("Insufficient funds")
		return
	}
	a.Balance -= amount
}
func (a Account) Display() {
	fmt.Printf("Tài khoản: %s, Số dư: %.2f\n", a.Name, a.Balance)
}

type Transaction interface {
	Process()
}

type Deposit struct {
	Account *Account
	Amount  float64
}

func (d Deposit) Process() {
	d.Account.Deposit(d.Amount)
}

type Withdraw struct {
	Account *Account
	Amount  float64
}

func (w Withdraw) Process() {
	w.Account.Withdraw(w.Amount)
}

// Hàm xử lý danh sách giao dịch
func ProcessTransactions(transactions []Transaction) {
	fmt.Println("Bắt đầu xử lý giao dịch:")
	for _, t := range transactions {
		t.Process()
	}
	fmt.Println("Kết thúc xử lý giao dịch.")
}

func main() {
	account := &Account{ID: 1, Name: "An", Balance: 1000.0}
	// Hiển thị thông tin ban đầu
	account.Display()
	transactions := []Transaction{
		Deposit{Account: account, Amount: 500.0},   // Nạp 500
		Withdraw{Account: account, Amount: 300.0},  // Rút 300
		Withdraw{Account: account, Amount: 1500.0}, // Rút thất bại
	}

	ProcessTransactions(transactions)
	account.Display()

}
