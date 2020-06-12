package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)  // 一个reflect.Type
	fmt.Println(t.String()) // int
	fmt.Println(t)          // int

	fmt.Println("==================================")

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File

	fmt.Println("==================================")

	fmt.Printf("%T\n", 3) // int

	fmt.Println("==================================")

	v := reflect.ValueOf(3) // 一个reflect.Value
	fmt.Println(v)          // 3
	fmt.Printf("%v\n", v)   // 3
	fmt.Println(v.String()) // <int Value>

	fmt.Println("==================================")

	t2 := v.Type()
	fmt.Println(t2.String()) // int

	fmt.Println("==================================")

	v2 := reflect.ValueOf(3) // a reflect.Value
	x2 := v2.Interface()     // an interface
	i2 := x2.(int)           // an int
	fmt.Printf("%d\n", i2)   // 3
}
