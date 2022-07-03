package main

import "fmt"

type student struct {
	name  string
	age   int
	class int
}

func (s student) Print() {
	fmt.Printf("Name: %v, Age: %v class: %v\n",
		s.name, s.age, s.class)
}

func (s *student) PrintPtr() {
	fmt.Printf("Name: %v, Age: %v class: %v\n",
		s.name, s.age, s.class)
}

func (s *student) Promote() {
	s.class++
}

func main() {
	fmt.Println("Go structs are awesome")
	var s1 student
	var s2 student
	var studentPtr *student

	s1 = student{name: "bala", age: 29, class: 5}
	s2 = student{"dhilip", 27, 6}
	studentPtr = &s2

	s1.Promote()
	s1.Print()

	s1.Promote()
	s2.Print()
	studentPtr.Print()
}
