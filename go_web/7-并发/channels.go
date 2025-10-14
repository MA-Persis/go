package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum

	fmt.Println("sum:", sum)
}

func main() {
	a := []int{1, 3, 5, 7, 9}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	x := <-c
	fmt.Println(x)
}
