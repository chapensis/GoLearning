package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func main() {
	var list1 = []string{"yang", "chang"}
	var list2 = []string{"chen", "jie", "ying"}
	fmt.Println(k(list1))
	fmt.Println(k(list2))

	Add(list1)
	Add(list1)
	Add(list1)

	Add(list2)
	Add(list2)

	fmt.Println(Count(list1))
	fmt.Println(Count(list2))
}
