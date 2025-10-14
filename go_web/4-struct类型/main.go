package main

import "fmt"

type person struct {
	name string
	age  int
}

type Student struct {
	person
	weight int
}

func Older(p1, p2 person) (person, string) {
	if p1.age > p2.age {
		return p1, p1.name
	}

	p2.name = "test"
	return p2, p2.name
}

func main() {
	tom := person{"tom", 25}
	fmt.Println(tom)

	jack := person{"jack", 25}

	p, name := Older(tom, jack)
	fmt.Println(p, name)

	mark := Student{person: person{"mark", 20}, weight: 110}
	fmt.Println(mark)
}
