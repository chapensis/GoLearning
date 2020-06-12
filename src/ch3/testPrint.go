package main

import (
	"bytes"
	"fmt"
	"math"
	"time"
)

func testNumberPrint() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := 0; i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	var temp uint = 1
	temp = temp - 1
	fmt.Println(temp)
}

func testNumberPrint2() {
	fmt.Println("===================testNumberPrint2=======================")
	o := 0666
	fmt.Printf("%d, %[1]o, %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d, %[1]x, %#[1]x, %#[1]X\n", x)
}

func testCharPrint() {
	fmt.Println("===================testCharPrint=======================")
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}

func testFloat() {
	fmt.Println("===================testFloat=======================")
	var f float32 = 1 << 24
	fmt.Println(f == (f + 1)) // 输出是true,因为已经丢失精度了

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func testNan() {
	fmt.Println("===================testNan=======================")
	nan := math.NaN()
	// 都是false,所以用==比较没有意义
	fmt.Println(nan == nan, nan < nan, nan > nan)
	// 可以用isNaN判断
	fmt.Println(math.IsNaN(nan))
}

func testString() {
	fmt.Println("===================testString=======================")
	str := "hello, 世界"
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	name := "杨昌小天才"
	fmt.Printf("%q\n", name[1])
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func testTimeDuration() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
}

func main() {
	testNumberPrint2()
	testCharPrint()
	testFloat()
	testNan()
	testString()

	fmt.Println(intsToString([]int{1, 2, 3}))

	testTimeDuration()
}
