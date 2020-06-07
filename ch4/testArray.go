package main

import (
	"crypto/sha256"
	"fmt"
)

func testArray01() {
	fmt.Println("==================testArray01======================")
	var a [3]int // 3个整数的数组
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d %d\n", 0, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2])
	fmt.Println(r[2])

	// 如果出现省略号，则数组长度由初始化数组的元素个数决定
	s := [...]int{1, 2, 3, 4}
	fmt.Println(len(s))
}

type Currency int
type MyValue int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

const (
	m1 MyValue = iota
	m2
)

type MyStruct struct {
	name string
	age  int
}

func testArray02() {
	fmt.Println("==================testArray2======================")
	symbol := [...]string{USD: "$", EUR: "#", GBP: "@", RMB: "!"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println(0, symbol[m1])
}

func testArray03(myarray [3]int) {
	fmt.Println("==================testArray03======================")
	myarray[0] = 100
}

func testStruct(myStruct MyStruct) {
	fmt.Println("==================testStruct======================")
	myStruct.age = 10
	myStruct.name = "chenjieying"
}

func testArray04() {
	fmt.Println("==================testArray04======================")
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // 长度不同不能编译
}

func testArray05() {
	fmt.Println("==================testArray05======================")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func main() {
	testArray01()

	testArray02()

	myarray := [3]int{1, 2, 3}
	fmt.Println(myarray[0])
	testArray03(myarray)
	fmt.Println(myarray[0])

	var myStruct = MyStruct{
		"yangchang",
		1}
	fmt.Println(myStruct.name)
	fmt.Println(myStruct.age)
	testStruct(myStruct)
	fmt.Println(myStruct.name)
	fmt.Println(myStruct.age)

	testArray04()

	testArray05()
}
