package main

import "fmt"

func main() {
	//equivalent to var cars [4]string as done in basic
	cars := [...]string{ //same as cars := [4]string
		"maruti",
		"huyndai",
		"toyota",
		"tata",
	}
	fmt.Printf("%q\n", cars)
	fmt.Printf("%#v\n", cars)
	fmt.Printf("length of cars array is %d\n", len(cars))
}

/* result

["maruti" "huyndai" "toyota" "tata"]
[4]string{"maruti", "huyndai", "toyota", "tata"}
length of cars array is 4

*/
