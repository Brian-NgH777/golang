package interfaces

import (
	"fmt"
)
// Interface
// Multiple Interface
// Embed Interface
// empty Interface

type NextAnimal interface {
	Animal
	Movement
}

type Movement interface {
	move()
	move2()
}

type Animal interface {
	speak()
}

type Dog struct {}
type Test struct {
	index int
}


func (d Dog) speak() { // implement Animal (method)
	fmt.Println("gau gau")
}

func (d Dog) move() { // implement Movement
	fmt.Println("4 chan")
}

func (d *Dog) move2() { // implement Movement
	fmt.Println("4 chan pointer")
}

func interface() {
	var animal Animal
	animal = Dog{}
	animal.speak()

	// Multiple Interface
	dog2 := Dog{}
	var m Movement = dog2
	m.move()
	var a Animal = dog2
	a.speak()

	// Embed Interface
	dog3 := Dog{}
	var na NextAnimal = dog3
	na.move()
	na.speak()

	// pointer
	dog4 := Dog{}
	var na1 NextAnimal2 = &dog4
	na1.move2()

	// empty Interface
	// bất cứ kiểu type nào trong golang đều tự động implement interface
	goout(1999)
	goout(22.333)
	goout("asdasd")
	d:= Test {
		index: 222
	}
	goout(d)


}

func goout(i interface{}) {
	fmt.Println("iiiii", i)
}