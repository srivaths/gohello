package main

import "fmt"

func main() {
	fmt.Printf("Enter temp in F: ")
	var f float64
	fmt.Scanf("%f", &f)
	c := ((f - 32) / 9) * 5
	fmt.Printf("%f degrees F  = %f degrees C\n", f, c)
}
