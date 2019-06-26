package main

import "fmt"

func main() {
	var (
		speed int
		ratio float32
		off   bool
		msg   string
	)

	//type should omit type bool from declaration of var safe; it will be inferred from the right-hand sidego
	var safe = true

	//or if we know the type then use shorthand declaration

	unsafe := false

	_ = unsafe //use blank identifier for making happy go compiler

	fmt.Println("speed is :", speed)
	fmt.Printf("Type of speed is %T\n", speed)

	fmt.Println("ratio is :", ratio)
	fmt.Printf("Type of ratio is %T\n", ratio)

	fmt.Println("off is :", off)
	fmt.Printf("Type of off is %T\n", off)

	fmt.Printf("msg is :%q\n", msg)
	fmt.Printf("Type of msg is %T\n", msg)

	fmt.Println("safe is :", safe)
	fmt.Printf("Type of safe is %T", safe)
}
