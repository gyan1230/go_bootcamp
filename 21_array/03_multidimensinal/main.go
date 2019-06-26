package main

import "fmt"

func main() {
	students := [...][3]float64{
		{1, 2, 3},
		{4, 5.5, 6},
	}
	var sum float64
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}
	l := float64(len(students) * len(students[0]))
	fmt.Printf("Avg is %.4f\n", sum/l)
}
