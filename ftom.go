package main

import "fmt"

func main() {
	var feet float64
	fmt.Printf("Provide length in feet: ")
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Printf("%f feet corresponds to %0.2f meters.\n", feet, meters)
}
