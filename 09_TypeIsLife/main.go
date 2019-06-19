package main

import "fmt"

type ounce float64
type gram float64

func main() {
	var o ounce
	var g gram
	g = 100
	o = ounce(g) * 0.03245
	fmt.Println(g, "gram is", o, "ounce")
}
