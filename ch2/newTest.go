package main

import "fmt"

func main()  {
	p := new(int)
	fmt.Println("*p:", *p)
	fmt.Println("p:", p)
	*p = 2
	fmt.Println("*p:", *p)
}
