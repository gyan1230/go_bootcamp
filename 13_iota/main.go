package main

import "fmt"

func main() {
	const (
		a = iota * 2 //0x2
		b            //1x2
		c            //2x2
		d            //3x2
		e            //4x2
	)
	fmt.Println(a, b, c, d, e)
}
