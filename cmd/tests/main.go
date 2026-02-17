package main

import "fmt"

func main() {
	res := Div(1, 2)
	fmt.Println(res)
}

func Div(a, b float64) float64 {
	return a / b
}
