package main

import (
	"fmt"
	"math"
)

func main() {
	//int8(math.MaxInt8 + 1)	//constant 128 overflows int8
	//max := math.MaxInt8		//shorthand no wrap around
	var max int8 = math.MaxInt8 //we have to use for wrap around
	var min int64 = math.MinInt64

	fmt.Println("max:	", max)
	//	max++
	fmt.Println("max+1: ", max+1)

	fmt.Println("min:   ", min)
	//	min++
	fmt.Println("min-1:  ", min-1)

	fmt.Println("equal?", math.MaxInt64 == min-1)

	/*
		OUTPUT:
			max:     127
			max++:  -128
			min:     -9223372036854775808
			min++:  9223372036854775807
			equal? true
	*/

}
