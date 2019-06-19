package main

import "fmt"

func main() {
	var val byte
	val = 0
	fmt.Printf("%08b : %d\n", val, val)
	val = 255
	fmt.Printf("%8b : %d\n", val, val)
}
