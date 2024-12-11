package main

import "fmt"

/*
	Bài tập 5: Hệ thống quản lý thú cưng
	Tạo interface Animal với các phương thức:
	Speak() string
	Eat(food string)
	Tạo các struct Dog và Cat, mỗi struct có các thuộc tính riêng như Name, Breed, HungerLevel.
	Viết hàm thay đổi trạng thái đói của thú cưng (sử dụng pointer).
	Triển khai một slice Animal để quản lý nhiều thú cưng và gọi phương thức Speak() của từng con.
*/
// Định nghĩa interface Animal
type Animal interface {
	Speak() string
	Eat(food string)
}

// Struct Dog đại diện cho chó
type Dog struct {
	Name        string
	Breed       string
	HungerLevel int
}

// Triển khai phương thức Speak cho Dog
func (d Dog) Speak() string {
	return fmt.Sprintf("%s (chó %s): Gâu gâu!", d.Name, d.Breed)
}

// Triển khai phương thức Eat cho Dog
func (d *Dog) Eat(food string) {
	fmt.Printf("%s đang ăn %s...\n", d.Name, food)
	d.HungerLevel -= 2
	if d.HungerLevel < 0 {
		d.HungerLevel = 0
	}
	fmt.Printf("Mức độ đói của %s giờ là %d.\n", d.Name, d.HungerLevel)
}

// Struct Cat đại diện cho mèo
type Cat struct {
	Name        string
	Breed       string
	HungerLevel int
}

// Triển khai phương thức Speak cho Cat
func (c Cat) Speak() string {
	return fmt.Sprintf("%s (mèo %s): Meo meo!", c.Name, c.Breed)
}

// Triển khai phương thức Eat cho Cat
func (c *Cat) Eat(food string) {
	fmt.Printf("%s đang ăn %s...\n", c.Name, food)
	c.HungerLevel -= 3
	if c.HungerLevel < 0 {
		c.HungerLevel = 0
	}
	fmt.Printf("Mức độ đói của %s giờ là %d.\n", c.Name, c.HungerLevel)
}

// Hàm chính để quản lý thú cưng
func main() {
	// Tạo các thú cưng
	dog := &Dog{Name: "Buddy", Breed: "Golden Retriever", HungerLevel: 5}
	cat := &Cat{Name: "Whiskers", Breed: "Siamese", HungerLevel: 7}

	// Quản lý thú cưng bằng slice Animal
	pets := []Animal{dog, cat}

	// Gọi phương thức Speak() cho từng thú cưng
	fmt.Println("Các thú cưng lên tiếng:")
	for _, pet := range pets {
		fmt.Println(pet.Speak())
	}

	// Gọi phương thức Eat() cho từng thú cưng
	fmt.Println("\nCho các thú cưng ăn:")
	for _, pet := range pets {
		pet.Eat("thức ăn cho thú cưng")
	}
}
