package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func testStruct01() {
	fmt.Println("====================testStruct01==========================")
	var dilbbert Employee
	dilbbert.Salary = dilbbert.Salary - 5000
	fmt.Println(dilbbert.Salary)
	position := &dilbbert.Position
	*position = "Senior " + *position
	fmt.Println(*position)

	var emploteeOfTheMonth *Employee = &dilbbert
	emploteeOfTheMonth.Position += " (proactive team player)"
	// 上面等价于下面语句
	(*emploteeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(emploteeOfTheMonth.Position)
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func testStruct02() {
	fmt.Println("====================testStruct02==========================")
	var w Wheel
	w.X = 9
	w.Circle.Point.X = 10

	w.Y = 8
	w.Circle.Point.Y = 10
	w.Radius = 5
	w.Spokes = 20

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Println(w)
	fmt.Printf("%#v\n", w)
}

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

func testStruct03() {
	fmt.Println("====================testStruct03==========================")
	var movies = []Movie{
		{Title: "Casablana1", Year: 1942, Color: false, Actors: []string{"Humphrey1", "Ingrid1"}},
		{Title: "Casablana2", Year: 1943, Color: true, Actors: []string{"Humphrey2", "Ingrid2"}},
		{Title: "Casablana3", Year: 1944, Color: false, Actors: []string{"Humphrey3", "Ingrid3"}},
		{Title: "Casablana4", Year: 1945, Color: false, Actors: []string{"Humphrey4", "Ingrid4"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marchaling failed: %s", err)
	}
	fmt.Printf("data: %s\n", data)

	data2, err2 := json.MarshalIndent(movies, "", "    ")
	if err2 != nil {
		log.Fatalf("JSON marchaling failed: %s", err2)
	}
	fmt.Printf("%s\n", data2)

	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshal failed: %s", err)
	}
	fmt.Println(titles)
	fmt.Printf("%v\n", titles)
}

func main() {
	testStruct01()
	testStruct02()
	testStruct03()
}
