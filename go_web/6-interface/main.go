package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, i am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("La, la la la la la la", lyrics)
}

func (h *Human) Guzzle(beer string) {
	fmt.Println("Guzzle", beer)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, i am %s, work on %s, you can call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beer string)
}

type YongChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {
	mike := Student{Human{"mike", 25, "13669257692"}, "MIT", 100}
	paul := Student{Human{"paul", 15, "13669257693"}, "MIT2", 100}
	sam := Employee{Human{"sam", 20, "13669257694"}, "Ali", 100}
	tom := Employee{Human{"tom", 10, "13669257692"}, "Ali2", 100}

	var i Men

	i = &mike
	fmt.Println("This is mike, a student")
	i.SayHi()
	i.Sing("mike")
	i.Guzzle("mike")

	i = &tom
	fmt.Println("This is tom, a student")
	i.SayHi()
	i.Sing("tom")

	x := make([]Men, 3)
	x[0], x[1], x[2] = &mike, &paul, &sam

	for _, value := range x {
		value.SayHi()
	}
}
