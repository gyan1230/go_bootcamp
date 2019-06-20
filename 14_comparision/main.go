package main

import "fmt"

func main() {
	fmt.Println("3 == 3", 3 == 3)
	fmt.Println(`"hi" != "HI" `, "hi" != "HI")
	fmt.Println("3 < 2.5 ", 3 < 2.5)   //it can compare bcoz untyped constant
	fmt.Println("4 <= 4.0 ", 4 <= 4.0) //it can compare bcoz untyped constant

	speed := 80
	fast := speed > 60
	slow := speed < 20

	fmt.Printf("type of fast is %T and slow is %T\n", fast, slow)
	fmt.Println("going fast?", fast)
	fmt.Println("going slow?", slow)
	fmt.Println("going at 80 speed?", speed == 80)

}
