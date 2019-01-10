package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(printIf(0))
	printSwitch(4)
	v1 := 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v4, v3, v2, v1)
}

// 这样会导致编译失败？
func printIf(x int) int {
	if x == 0 {
		return 5
	} else {
		return 0
	}
}

func printSwitch(i int) {
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4 5 6")
	default:
		fmt.Println("Default")
	}

	printFor()
}

func printFor() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	a := 3
	b := 4
	a, b = b, a
	fmt.Println(a, ",", b)

	fmt.Println(Add(3, 3))
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}

	myFunc(1, 2, 3, 4, 5)

	return a + b, nil
}

// 不定参数的传入
func myFunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	myfunc2(args...)
	myfunc2(args[1:]...)
}

func myfunc2(args ...int) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
}

// 任意类型的不定参数
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknow type.")
		}
	}

	printDemo()
}

func printDemo() {
	// 匿名函数
	f := func(x, y int) int {
		return x + y
	}
	result := f(1, 2)
	fmt.Println(result)

	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

	printBBox()
}

func printBBox() {

	// 闭包
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()

	j *= 2
	a()
}
