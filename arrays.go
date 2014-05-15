package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}
	fmt.Println("x is ", x)

	var y [3]int = [3]int{4, 5, 6}
	fmt.Println("y is ", y)

	z := [3]int{7,
		8,
		9}
	fmt.Println("z is ", z)
}
