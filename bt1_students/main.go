package main

import "fmt"

/*
Bài tập 1: Quản lý sinh viên
Tạo một struct Student với các thuộc tính: ID, Name, Age, Scores (danh sách điểm).
Viết một hàm để:
Thêm điểm vào danh sách Scores thông qua pointer.
Tính điểm trung bình của một Student.
Sử dụng slice để quản lý danh sách sinh viên và tìm sinh viên có điểm trung bình cao nhất.
*/
type Student struct {
	ID     int
	Name   string
	Age    int
	Scores []float64
}

func (s Student) AddScore(score float64) {
	s.Scores = append(s.Scores, score)
}
func (s Student) AverageScore() float64 {
	sum := 0.0
	for _, score := range s.Scores {
		sum += score
	}
	if len(s.Scores) == 0 {
		return 0.0
	}
	return float64(sum) / float64(len(s.Scores))
}
func FindTopStudent(students []Student) Student {
	if len(students) == 0 {
		return Student{}
	}
	topStudent := students[0]
	highestAverage := topStudent.AverageScore()
	for _, student := range students {
		avg := student.AverageScore()
		if avg > highestAverage {
			topStudent = student
			highestAverage = avg
		}
	}

	return topStudent
}
func main() {
	students := []Student{
		{ID: 1, Name: "An", Age: 20, Scores: []float64{8.0, 9.0}},
		{ID: 2, Name: "Bình", Age: 21, Scores: []float64{7.5, 8.5, 9.0}},
		{ID: 3, Name: "Châu", Age: 22, Scores: []float64{9.0, 9.5, 10.0}},
	}
	students[0].AddScore(10.0)
	// Hiển thị thông tin sinh viên và điểm trung bình
	fmt.Println("Danh sách sinh viên:")
	for _, student := range students {
		fmt.Printf("ID: %d, Tên: %s, Tuổi: %d, Điểm TB: %.2f\n",
			student.ID, student.Name, student.Age, student.AverageScore())
	}

	// Tìm sinh viên có điểm trung bình cao nhất
	topStudent := FindTopStudent(students)
	fmt.Printf("\nSinh viên có điểm TB cao nhất: %s (%.2f)\n", topStudent.Name, topStudent.AverageScore())
}
