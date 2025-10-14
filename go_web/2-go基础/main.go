package main

import "fmt"

const (
	i      = 100
	pi     = 3.1415
	prefix = "go"
)

const (
	x = iota
	y
	z
)

var (
	i_int         int
	pi_float32    float32
	prefix_string string
)

func main() {
	fmt.Println(y)

	varint1, varint2, varint3 := 1, 2, 3
	fmt.Println("varint:", varint1, varint2, varint3)

	const Pi = 3.1415926
	var pi float32 = Pi
	// pi2 := float32(Pi)
	fmt.Println(pi)

	var s string = "hello"
	// s[0] = 'c'
	c := []byte(s)
	c[0] = 's'
	fmt.Println(s)
	fmt.Printf("%s\n", c)

	s2 := string(c)
	fmt.Printf("%s\n", s2)

	m := "world"
	n := s + " " + m
	fmt.Printf("%s\n", n)

	k := `hello
world`
	fmt.Printf("%s\n", k)

	var arr [10]int
	arr[0] = 12
	arr[1] = 13
	fmt.Println(arr)

	arr_a := [3]int{1, 2, 3}
	arr_b := [10]int{1, 2, 3}
	arr_c := [...]int{4, 5, 6}
	fmt.Println(arr_a, arr_b, arr_c)

	double_arr := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7}}
	fmt.Println(double_arr)

	double_arr2 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7}}
	fmt.Println(double_arr2)

	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice, len(slice), cap(slice))

	slice_a := slice[:2]
	slice_b := slice[2:]
	fmt.Println(slice_a, slice_b, len(slice_a), cap(slice_b))

	slice_c := append(slice, 'e')
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice_a, len(slice_a), cap(slice_a))
	fmt.Println(slice_b, len(slice_b), cap(slice_b))
	fmt.Println(slice_c, len(slice_c), cap(slice_c))

	slice_a[0] = 'f'
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice_a, len(slice_a), cap(slice_a))
	fmt.Println(slice_b, len(slice_b), cap(slice_b))
	fmt.Println(slice_c, len(slice_c), cap(slice_c))
}
