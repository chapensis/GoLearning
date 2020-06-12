package main

import "fmt"

func testSlice01() {
	fmt.Println("======================testSlice01=======================")
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println("Q2 Length:", len(Q2), " Q2 Cap:", cap(Q2), Q2)
	fmt.Println("summer Length:", len(summer), " summer Cap:", cap(summer), summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// 新的slice会扩展原slice
	endlessSumber := summer[:5]
	fmt.Println(endlessSumber)
}

func reverse(mySlice []int) {
	for i, j := 0, len(mySlice)-1; i < j; i, j = i+1, j-1 {
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
	}
}

func testSlice02() {
	fmt.Println("======================testSlice02=======================")
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func testSliceAppend() {
	var runes []rune
	for _, r := range "Hello world" {
		runes = append(runes, r)
		fmt.Println(len(runes), cap(runes))
	}
	fmt.Printf("%q \n", runes)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func testSliceAppendInt() {
	fmt.Println("=====================testSliceAppendInt========================")
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func testSlice03() {
	fmt.Println("=====================testSlice03========================")
	data := []string{"one","", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func main() {
	testSlice01()
	testSlice02()

	testSliceAppend()

	testSliceAppendInt()

	testSlice03()
}
