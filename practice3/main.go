package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	Age  int
}
type studentList struct {
	students []student
}

func (sl *studentList) addStudent(s *student) {
	sl.students = append(sl.students, *s)
}
func (s *student) updateName(n string) {
	s.name = n
}
func (sl *studentList) updateNameSlist(s string) {
	for i := range sl.students {
		if strings.HasPrefix(strings.ToLower(s), strings.ToLower(sl.students[i].name)) {
			sl.students[i].updateName(s)
		}
	}
}
func main() {
	student1 := &student{
		name: "shiva",
		Age:  25,
	}
	var sList studentList
	sList.addStudent(student1)
	for i := range sList.students {
		if sList.students[i].name == "shiva" {
			sList.students[i].updateName("shiva KUMAR")
			fmt.Println(sList.students[i])
		}
	}
	fmt.Println(sList)
	sList.updateNameSlist("SHIVA Kumar Rayapuram")
	fmt.Println(sList)

}
