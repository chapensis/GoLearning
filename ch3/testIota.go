package main

import (
	"fmt"
	"math"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBoardcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main()  {
	fmt.Println(Saturday)

	fmt.Println(FlagUp)
	fmt.Println(FlagBoardcast)
	fmt.Println(FlagLoopback)
	fmt.Println(FlagPointToPoint)
	fmt.Println(FlagMulticast)

	fmt.Println(math.Pi)
}
