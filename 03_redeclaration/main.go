package main

import "fmt"

func main() {
	name := "gyan"

	//redeclaration name variable possible due to age new variable
	name, age := "Albert", 1857
	fmt.Println(name, age)

}
