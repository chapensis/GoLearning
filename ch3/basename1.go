package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		// 只取最后一个斜杠后的内容
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个点之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		// 只取最后一个斜杠后的内容
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))

	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename2("c.d.go"))
	fmt.Println(basename2("abc"))
}
