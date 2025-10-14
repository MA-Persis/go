package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	weight, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.weight * r.weight
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Human struct {
	name   string
	weight int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayhi() {
	fmt.Println("human", h.name)
}

func (e *Employee) sayhi() {
	fmt.Println("employee", e.Human.name)
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println(r1.area())
	fmt.Println(r2.area())
	fmt.Println(c1.area())
	fmt.Println(c2.area())

	mark := Student{Human{"mark", 25}, "school"}
	jack := Employee{Human{"jack", 25}, "company"}
	mark.sayhi()
	jack.sayhi()
}
