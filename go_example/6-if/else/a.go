// if 和 else 分支结构在 Go 中当然是直接了当的了。

package main

func main() {
	if 7%2 == 0 { // 这里是一个基本的例子。
		println("7 is even")
	} else {
		println("7 is odd")
	}

	if 8%4 == 0 { // 你可以不要 else 只用 if 语句。
		println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // 在条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用。
		println(num, "is negative")
	} else if num < 10 { // 注意，在 Go 中，你可以不适用圆括号，但是花括号是需要的。
		println(num, "has one digit")
	} else {
		println(num, "has multiple digits")
	}

	// Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句。
}
