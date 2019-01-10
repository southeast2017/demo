package main

import (
	"fmt"
)

// Go语言的变量声明，类型在变量之后
var v1 int
var v2 string
var v3 [10]int // 数组
var v4 []int   // 数组切片
var v5 struct {
	f int
}
var v6 *int           // 指针
var v7 map[string]int // map, key为string类型，value为int类型
var v8 func(a int) int

var (
	v9  int
	v10 string
)

var v11 int = 10
var v12 = 10

const a, b, c = 3, 4, "foo"

const (
	c0 = iota
	c1
	c2
)

const (
	Sunday = iota
	Monday
	Tuesdasy
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

// Go  		->  	C++
// int8		->		char
// uint8	->      unsigned char
// int16    ->      int
// uint16   ->      unsigned int
// int32    ->      long
// uint32   ->      unsigned long
// int64    ->      long long
// uint64   ->      unsigned long long

func main() {
	c := 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	printIota()
	printComplex()
	printString()
	printStringMap()
	printFunc()
}

func printIota() {
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
}

func printComplex() {
	var value1 complex64

	value1 = 3.2 + 12i
	value2 := 3.2 + 12i
	value3 := complex(3.2, 12)

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
}

// 字符串的内容不能在初始化后进行修改
func printString() {
	var str string

	str = "Hello world"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c. \n", str, ch)
	fmt.Println(str[0])
}

func printStringMap() {
	str := "Hello, world"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
}

func printFunc() {
	printStringUnicode()
	printArray()

	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println(array)

	printSliceArray()
	printSliceFunc()
}

func printStringUnicode() {
	// 每个字符的类型是rune
	str := "Hello, world"
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func printArray() {
	var array [20]int

	for i := 0; i < len(array); i++ {
		fmt.Println(i, array[i])
	}

	for i, v := range array {
		fmt.Println("Array element[", i, "]=", v)
	}
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println(array)
}

func printSliceArray() {

	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mySlice := myArray[:5]

	fmt.Println("Elements of myArray: ")

	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
}

func printSliceFunc() {
	printSliceDemo()
}

func printSliceDemo() {
	mySlice2 := make([]int, 5, 10)

	fmt.Println()
	for _, v := range mySlice2 {
		fmt.Print(v)
	}
	fmt.Println(len(mySlice2))

	fmt.Println(cap(mySlice2))

	mySlice2 = append(mySlice2, 1, 2, 3, 4, 5, 6)
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println(len(mySlice2))

	fmt.Println(cap(mySlice2))

	mySlice3 := make([]int, 3, 10)
	fmt.Println(len(mySlice3))
	mySlice3 = append(mySlice3, 2, 3, 4)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))

	mySlice2 = append(mySlice2, mySlice3...)
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	printSliceInSlice()
}

func printSliceInSlice() {
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]

	for _, v := range newSlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	printSliceCopy()
}

func printSliceCopy() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1中

	fmt.Println(slice1)
	fmt.Println(slice2)

	printMap()
}

func printMap() {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}

	// 变量的声明， personDB 声明是map类型的，string 是键的类型 PersonInfo是所放值的类型
	var personDB map[string]PersonInfo

	// 创建
	personDB = make(map[string]PersonInfo)

	// 元素赋值
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	person, ok := personDB["12345"]
	fmt.Println(person)

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234")
	}

	// 元素删除
	delete(personDB, "12345")
}
