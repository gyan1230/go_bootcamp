package main

import "fmt"

func main() {
	const (
		x       = 12
		y       = 3.4
		z int32 = 66
	)
	const min = 254
	var bb byte = min //here back bb = byte(min) is happening
	fmt.Printf("%8b:%T\n", bb, bb)

	ff := 1 + z
	fmt.Printf("%v : %T\n", ff, ff)
	var i = y
	var f float32 = x
	var b byte = x
	var r rune = x

	fmt.Printf("%v:%T\n%v:%T\n%v:%T\n%v:%T\n", i, i, f, f, b, b, r, r)
}
