package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func testMap01() {
	myMap := make(map[string]int)
	fmt.Println(myMap["yangchang"])
	myMap["yangchang"] = 28
	myMap["chenjieying"] = 29
	myMap["liyang"] = 27
	myMap["liuxianliang"] = 26
	for name, age := range myMap {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range myMap {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Println("================after sort================")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, myMap[name])
	}

	fmt.Println("================map init================")

	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages))
	fmt.Println(ages["yangchang"])
	// 以上操作都成功，但是下述操作会失败，因为没有初始化,可以执行 ages = make(map[string]int) 来初始化
	//ages["yangchang"] = 25

	name := "yangchang"
	age, ok := ages["yangchang"]
	if ok {
		fmt.Printf("name: %s is in map, age is %d\n", name, age)
	} else {
		fmt.Printf("name: %s is not in map, no age\n", name)
	}

	if age, ok := ages["yangchang"]; ok {
		fmt.Printf("name: %s is in map, age is %d\n", name, age)
	} else {
		fmt.Printf("name: %s is not in map, no age\n", name)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, vx := range x {
		if vy, ok := y[k]; !ok || vx != vy {
			return false
		}
	}
	return true
}

func testMap02() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "end" {
			break
		}
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	}
}

func main() {
	testMap01()

	testMap02()
}
