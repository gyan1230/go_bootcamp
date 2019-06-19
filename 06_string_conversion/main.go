package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	cel, _ := strconv.ParseFloat(arg, 64)
	far := (cel * 9 / 5) + 32
	fmt.Printf("%g celcius is %g farenhite", cel, far)
}
