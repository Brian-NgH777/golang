package methods

import "fmt"

type Student struct {
	name string
}

// non - struct

type String string

func (s String) append(str string) string {
	s = str
}

// define Method
// func (t Type) MethodName(params) returnType {}

// (t Type) => Receiver
// 1. value Receiver
// 2. pointer Receiver

func (s Student) changeName() string {
	s.name = "brian"
}

func (s *Student) changeName2() string {
	s.name = "brian"
}

func Method() {
	student := Student{name: "aaaa"}
	student.changeName()
	fmt.Println(student.name) // nó không có thay đổi vì nó copy cái student để đưa vào changeName() (value Receiver)
	// pointer
	student1 := &Student{name: "aaaa"}
	student1.changeName2()
	fmt.Println(student.name) // nó thay đổi

	// non - struct

	s1 := String("aaaa")
	newStr := s1.append("brainnnnnn")

	fmt.Println(newStr)

}
