package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

/*
	Bài 7: Quản lý tài khoản ngân hàng
	1. Tạo interface Account với các phương thức:
	Deposit(amount float64): Nạp tiền.
	Withdraw(amount float64) error: Rút tiền (kiểm tra số dư).
	GetBalance() float64: Lấy số dư.
	2. Tạo hai loại tài khoản:
	SavingsAccount với thuộc tính Balance và lãi suất (InterestRate).
	CheckingAccount với thuộc tính Balance và phí giao dịch (TransactionFee).
	3. Viết chương trình quản lý nhiều tài khoản trong slice []Account và thực hiện các giao dịch như nạp, rút tiền, và kiểm tra số dư.

	Bonus:
	- Thêm phương thức ApplyInterest() cho tài khoản tiết kiệm để tính lãi.
	- Sử dụng JSON hoặc cơ sở dữ liệu để lưu trạng thái các tài khoản.
	- Thêm tài khoản doanh nghiệp (BusinessAccount) với các quy định đặc biệt như hạn mức tối thiểu.
*/

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
}

// Struct SavingsAccount đại diện cho tài khoản tiết kiệm
type SavingsAccount struct {
	Balance      float64
	InterestRate float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
	fmt.Printf("Nạp %.2f vào tài khoản tiết kiệm. Số dư hiện tại: %.2f\n", amount, s.Balance)
}
func (s *SavingsAccount) Withdraw(amount float64) error {
	if s.Balance < amount {
		return errors.New("số dư không đủ để rút tiền")
	}
	s.Balance -= amount
	fmt.Printf("Rút %.2f từ tài khoản tiết kiệm. Số dư hiện tại: %.2f\n", amount, s.Balance)
	return nil
}
func (s *SavingsAccount) GetBalance() float64 {
	return s.Balance
}
func (s *SavingsAccount) ApplyInterest() {
	interest := s.Balance * s.InterestRate
	s.Balance += interest
	fmt.Printf("Lãi %.2f đã được cộng vào tài khoản tiết kiệm. Số dư mới: %.2f\n", interest, s.Balance)
}

// Struct CheckingAccount đại diện cho tài khoản vãng lai
type CheckingAccount struct {
	Balance        float64
	TransactionFee float64
}

func (c *CheckingAccount) Deposit(amount float64) {
	c.Balance += amount
	fmt.Printf("Nạp %.2f vào tài khoản vãng lai. Số dư hiện tại: %.2f\n", amount, c.Balance)
}
func (c *CheckingAccount) Withdraw(amount float64) error {
	totalAmount := amount + c.TransactionFee
	if totalAmount > c.Balance {
		return errors.New("số dư không đủ để rút tiền sau khi tính phí giao dịch")
	}
	c.Balance -= totalAmount
	fmt.Printf("Rút %.2f từ tài khoản vãng lai (phí: %.2f). Số dư hiện tại: %.2f\n", amount, c.TransactionFee, c.Balance)
	return nil
}
func (c *CheckingAccount) GetBalance() float64 {
	return c.Balance
}

func SaveAccountsToFile(accounts []Account, filename string) error {
	var data []map[string]interface{}
	for _, account := range accounts {
		switch acc := account.(type) {
		case *SavingsAccount:
			data = append(data, map[string]interface{}{
				"Type":         "SavingsAccount",
				"Balance":      acc.Balance,
				"InterestRate": acc.InterestRate,
			})
		case *CheckingAccount:
			data = append(data, map[string]interface{}{
				"Type":           "CheckingAccount",
				"Balance":        acc.Balance,
				"TransactionFee": acc.TransactionFee,
			})
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}
func LoadAccountsFromFile(filename string) ([]Account, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	var accounts []Account
	for _, item := range data {
		switch item["Type"] {
		case "SavingsAccount":
			accounts = append(accounts, &SavingsAccount{
				Balance:      item["Balance"].(float64),
				InterestRate: item["InterestRate"].(float64),
			})
		case "CheckingAccount":
			accounts = append(accounts, &CheckingAccount{
				Balance:        item["Balance"].(float64),
				TransactionFee: item["TransactionFee"].(float64),
			})
		}
	}
	return accounts, nil
}

func main() {
	// Tạo danh sách tài khoản
	accounts := []Account{
		&SavingsAccount{Balance: 1000, InterestRate: 0.02},
		&CheckingAccount{Balance: 500, TransactionFee: 5},
	}
	// Thực hiện các giao dịch

	fmt.Println("Thực hiện giao dịch:")
	accounts[0].Deposit(200) // Nạp tiền vào tài khoản tiết kiệm
	accounts[1].Deposit(100) // Nạp tiền vào tài khoản vãng lai
	if savings, ok := accounts[0].(*SavingsAccount); ok {
		savings.ApplyInterest()
	}
	err := accounts[0].Withdraw(500) // Rút tiền từ tài khoản tiết kiệm

	if err != nil {
		fmt.Println("Lỗi:", err)
	}
	fmt.Println(accounts[0])
	err = accounts[1].Withdraw(550) // Rút tiền từ tài khoản vãng lai (phải tính phí)

	if err != nil {
		fmt.Println("Lỗi:", err)
	}
	// Lưu tài khoản vào file
	err = SaveAccountsToFile(accounts, "accounts.json")
	if err != nil {
		fmt.Println("Lỗi khi lưu tài khoản:", err)
	}

	// Đọc tài khoản từ file
	loadedAccounts, err := LoadAccountsFromFile("accounts.json")
	if err != nil {
		fmt.Println("Lỗi khi đọc tài khoản:", err)
	} else {
		fmt.Println("\nTài khoản đã được tải lại từ file:")
		for i, account := range loadedAccounts {
			fmt.Printf("Tài khoản %d: Số dư = %.2f\n", i+1, account.GetBalance())
		}
	}

}
