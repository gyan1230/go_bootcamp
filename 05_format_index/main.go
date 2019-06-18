package main

import "fmt"

func main() {
	var (
		age  = 12
		name = "xxx"
	)
	fmt.Printf("Age is :%d and name is :%s\n", age, name)

	//formatting index
	fmt.Printf("Age is :%[2]v and name is :%[1]v\n", age, name)
}
