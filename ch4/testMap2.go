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

var graph = make(map[string]map[string]bool)

func addEdge(from, to string)  {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string)  bool{
	return graph[from][to]
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

	fmt.Println(hasEdge("yang", "chang"))
}
