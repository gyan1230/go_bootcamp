package main

import "fmt"

func main() {
	const (
		x = 12
		y = 3.4
	)
	//Valid operation in mismatch type due to untyped constant
	f := x * y
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	var (
		a = 12
		b = 3.2
	)
	//	m := a * b //invalid operation type mismatch in variable
	m := a * int(b)
	fmt.Println("int conversion", m)
	//OR
	n := float64(a) * b
	fmt.Printf("float conversion :%g - %T", n, n)
}
