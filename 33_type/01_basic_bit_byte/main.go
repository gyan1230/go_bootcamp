package main

import "fmt"

func main() {
	fmt.Printf("%b\n", 2)
	fmt.Printf("%02b\n", 0)
	fmt.Printf("%06b\n", 4)
	fmt.Printf("%08b\n", 254)

	var b byte
	fmt.Printf("%08b\n", b)
	b = 255
	fmt.Printf("%08b\n", b)
}
