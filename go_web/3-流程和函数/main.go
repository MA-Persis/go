package main

import "fmt"

type testInt func(int) bool

func computevalue() int32 {
	var id int32 = 5
	return id
}

func myfunc() {
	i := 0
Here:
	println(i)
	i++
	if i <= 5 {
		goto Here
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reurn_a_b(a, b int) (int, int) {
	return a, b
}

func sum(a, b int) (add int, product int) {
	add = a + b
	product = a * b
	return
}

func myfunc2(arg ...int) int {
	fmt.Println(arg)

	total := 0
	for _, v := range arg {
		total += v
	}
	return total
}

func add(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	if x := computevalue(); x > 0 {
		fmt.Println("above")
	} else {
		fmt.Println("below")
	}

	myfunc()

	x := 2
	y := 3

	fmt.Println(max(x, y))

	a, b := reurn_a_b(x, y)
	fmt.Println(a, b)

	c, d := sum(x, y)
	fmt.Println(c, d)

	fmt.Println(myfunc2(1, 2, 3))

	e := []int{2, 3, 4}
	fmt.Println(myfunc2(e...))

	f := 6
	fmt.Println(add(&f))

}
