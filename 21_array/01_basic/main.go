package main

import "fmt"

const (
	single = 1
	team   = 3
	total  = single + team
)

func main() {
	var name [total]string
	fmt.Printf("%T\n", name)
	fmt.Printf("%q\n", name)
	fmt.Printf("%#v\n", name)
	var (
		s [single]string
		t [team]string
	)
	name[0] = "GYan"
	name[1] = "Tom"
	name[2] = "Dick"
	name[3] = "Harry"

	s[0] = name[0]
	for i := range t {
		t[i] = name[i+1]
	}

	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", t)

	var (
		active [total]bool
	)
	active[0] = true
	active[len(name)-1] = true
	fmt.Println("Active names are:")
	for i, ok := range active {
		if ok {
			fmt.Printf("%q\n", name[i])
		}
	}

}
