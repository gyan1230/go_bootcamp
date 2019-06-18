package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	sum := 0
	for i, v := range os.Args {
		fmt.Println(i, v)
		sum++
	}
	fmt.Println("Total argument is :", sum-1)
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
