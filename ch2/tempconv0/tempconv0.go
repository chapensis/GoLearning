package main

import "fmt"
import "../tempconv"

func main() {
	tempC := tempconv.Celsius(50)
	tempF := tempconv.CToF(tempC)
	fmt.Println(tempF)

	tempC = tempconv.FToC(tempF)
	fmt.Println(tempC)

	fmt.Printf("%g\n", tempconv.BoilingC - tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF - tempconv.CToF(tempconv.BoilingC))
	//fmt.Printf("%g\n", boilingF - FreezingC) // 编译错误：类型不匹配

	fmt.Println(tempC.String())
}
