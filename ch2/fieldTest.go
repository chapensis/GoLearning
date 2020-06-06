package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
}

func test2() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()
}

var cwd string

func test3() {
	/*
		// 编译错误，未使用cwd
		cwd, err  := os.Getwd()
		if err != nil {
			log.Fatalf("os.Getwd failed: %v", err)
		}
	*/
}

func test4() {
	// 优先会自行创建一个cwd
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v\n", err)
	}
	// 使用的还是内部的cwd
	log.Printf("Working Directory = %s", cwd)
}

func main() {
	test2()

	test4()

	log.Printf("Working Directory = %s", cwd)
}
