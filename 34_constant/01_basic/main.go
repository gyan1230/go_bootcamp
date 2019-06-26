package main

import "fmt"

func main() {
	//Typed constant is having issue
	// const (
	// 	min int     = 1
	// 	max float32 = min * 3.2 //cannot use min * 3.2 (type int) as type float32 in const initializer
	// )

	//Untyped constant is a kind
	const (
		min = 1 + 1     // int type
		max = min * 3.2 // float64 type
	)

	fmt.Println("min :", min, "\n", "max : ", max)
	fmt.Printf("min is %T\nmax is %T\n", min, max)

	/*
		min is int
		max is float64
	*/

	const untyped = 11

	var i int = untyped
	var f float64 = untyped // behind var f = float64(untyped)
	var b byte = untyped
	var j int32 = untyped
	var r rune = untyped
	//cannot use untyped (type int) as type string in assignment
	//var s string = untyped

	fmt.Printf("Types are i : %T and value is %d\n", i, i)
	fmt.Printf("Types are f : %T and value is %f\n", f, f)
	fmt.Printf("Types are b : %T and value is %d\n", b, b)
	fmt.Printf("Types are j : %T and value is %d\n", j, j)
	fmt.Printf("Types are r : %T and value is %d\n", r, r)

	/*
		Types are i : int and value is 11
		Types are f : float64 and value is 11.000000
		Types are b : uint8 and value is 11
		Types are j : int32 and value is 11
		Types are r : int32 and value is 11
	*/

}
