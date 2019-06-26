package main

import "fmt"

func main() {
	var (
		myAge   = 32
		yourAge = 61
		average float32
	)
	average = float32(myAge+yourAge) / 2
	fmt.Printf("Average is :%.2f\n", average)
	fmt.Println("Average is ", average)
}
