package main

import "fmt"

/*
	Bài 6: Hệ thống quản lý nhân viên
	1. Tạo interface Employee với các phương thức:
	CalculateSalary() float64: Tính lương.
	DisplayInfo(): Hiển thị thông tin nhân viên.
	2. Tạo hai struct:
	FullTimeEmployee với các thuộc tính như Name, BaseSalary, Bonus.
	PartTimeEmployee với các thuộc tính như Name, HourlyRate, HoursWorked.
	3. Triển khai interface Employee cho hai loại nhân viên.
	4. Viết hàm nhận danh sách Employee và tính tổng lương phải trả cho tất cả nhân viên.
*/

type Employee interface {
	CalculateSalary() float64
	DisplayInfo()
}
type FullTimeEmployee struct {
	Name       string
	BaseSalary float64
	Bonus      float64
}

func (f FullTimeEmployee) CalculateSalary() float64 {
	return f.BaseSalary + f.Bonus
}
func (f FullTimeEmployee) DisplayInfo() {
	fmt.Printf("Full time employee %s earns %.2f per hour.\n", f.Name, f.CalculateSalary())
}

type PartTimeEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked int
}

func (p PartTimeEmployee) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}
func (p PartTimeEmployee) DisplayInfo() {
	fmt.Printf("Part time employee %s earns %.2f per hour.\n", p.Name, p.CalculateSalary())
}
func totalSalary(employees []Employee) float64 {
	total := 0.0
	for _, e := range employees {
		total += e.CalculateSalary()
	}
	return total
}
func main() {
	employees := []Employee{
		FullTimeEmployee{Name: "John", BaseSalary: 10000, Bonus: 4000},
		FullTimeEmployee{Name: "Caster", BaseSalary: 40000, Bonus: 5000},
		PartTimeEmployee{Name: "Sala", HourlyRate: 30, HoursWorked: 50},
		PartTimeEmployee{Name: "Mary", HourlyRate: 20, HoursWorked: 40},
	}
	for _, e := range employees {
		e.DisplayInfo()
	}
	fmt.Printf("Total salary: %.2f\n", totalSalary(employees))
}
