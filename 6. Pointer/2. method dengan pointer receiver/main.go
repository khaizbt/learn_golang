package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {
	student := Student{1, "Khaiz Badaru Tammam", 3.49}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}

func (student *Student) graduate() {
	student.Name = student.Name + ",S.Kom"
}
