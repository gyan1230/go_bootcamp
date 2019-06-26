package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	raw := `hi
	It's a raw string
	literal in multi line`

	normal := "normal string"
	chinese := "Hi, 世界"

	fmt.Printf("Raw string literal type is %T and value is %q\n", raw, raw)
	fmt.Printf("Normal string literal type is %T and value is %q\n", normal, normal)

	//String concatination
	fmt.Println("Gyan" + "chand" + " " + "mohanty")

	//len function returns no of bytes in string value
	fmt.Println("length of normal string is:", len(normal))
	fmt.Println("length of chinese string is:", len(chinese))
	fmt.Println("No of characters in chinese string is:", utf8.RuneCountInString(chinese))

	fmt.Printf("Int(255) convert to string: %q", strconv.Itoa(255))

}
