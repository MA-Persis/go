// Go 支持在结构体类型中定义方法 。

package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // 这里的 area 方法有一个接收器类型 rect。
	return r.height * r.width
}

func (r rect) perim() int { // 可以为值类型或者指针类型的接收器定义方法。这里是一个值类型接收器的例子。
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area()) // 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("perim:", r.perim())

	rp := &r // Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
