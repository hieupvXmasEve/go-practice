package main

import (
	"fmt"
)

/*
	Bài 4: Hệ thống quản lý đội thể thao

	1.	Tạo interface Player với các phương thức:
	•	Play() string: Mô tả hành động khi chơi.
	•	Train(hours int): Tăng cường kỹ năng dựa trên số giờ luyện tập.
	2.	Tạo ba loại cầu thủ:
	•	FootballPlayer với thuộc tính Name, Position, SkillLevel.
	•	BasketballPlayer với thuộc tính Name, Height, SkillLevel.
	•	TennisPlayer với thuộc tính Name, Ranking, SkillLevel.
	3.	Viết chương trình:
	•	Quản lý danh sách []Player cho một đội thể thao.
	•	Cho các cầu thủ luyện tập và hiển thị kỹ năng của họ sau khi luyện tập.
*/

type Player interface {
	Play() string
	Train(hours int)
	GetSkillLevel() int
}

type FootballPlayer struct {
	Name       string
	Position   string
	SkillLevel int
}

func (f *FootballPlayer) Play() string {
	return fmt.Sprintf("%s chơi ở vị trí %s.", f.Name, f.Position)
}

func (f *FootballPlayer) Train(hours int) {
	f.SkillLevel += hours
}

func (f *FootballPlayer) GetSkillLevel() int {
	return f.SkillLevel
}

type BasketballPlayer struct {
	Name       string
	Height     float64
	SkillLevel int
}

func (b *BasketballPlayer) Play() string {
	return fmt.Sprintf("%s chơi bóng rổ với chiều cao %.2f m.", b.Name, b.Height)
}

func (b *BasketballPlayer) Train(hours int) {
	b.SkillLevel += hours * 2 // giả sử mỗi giờ luyện tập tăng kỹ năng gấp đôi
}

func (b *BasketballPlayer) GetSkillLevel() int {
	return b.SkillLevel
}

type TennisPlayer struct {
	Name       string
	Ranking    int
	SkillLevel int
}

func (t *TennisPlayer) Play() string {
	return fmt.Sprintf("%s chơi quần vợt với thứ hạng %d.", t.Name, t.Ranking)
}

func (t *TennisPlayer) Train(hours int) {
	t.SkillLevel += hours / 2 // giả sử mỗi 2 giờ luyện tập tăng kỹ năng 1 điểm
}

func (t *TennisPlayer) GetSkillLevel() int {
	return t.SkillLevel
}

func main() {
	// Tạo danh sách cầu thủ
	team := []Player{
		&FootballPlayer{Name: "David", Position: "Tiền đạo", SkillLevel: 50},
		&BasketballPlayer{Name: "Michael", Height: 1.98, SkillLevel: 70},
		&TennisPlayer{Name: "Serena", Ranking: 1, SkillLevel: 90},
	}

	// Hiển thị thông tin ban đầu
	fmt.Println("Thông tin cầu thủ ban đầu:")
	for _, player := range team {
		fmt.Println(player.Play(), "Kỹ năng:", player.GetSkillLevel())
	}

	// Tất cả cầu thủ luyện tập
	fmt.Println("\nCầu thủ luyện tập 5 giờ:")
	for _, player := range team {
		player.Train(5)
	}

	// Hiển thị thông tin sau khi luyện tập
	fmt.Println("\nThông tin cầu thủ sau luyện tập:")
	for _, player := range team {
		fmt.Println(player.Play(), "Kỹ năng:", player.GetSkillLevel())
	}
}
