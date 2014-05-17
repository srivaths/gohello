package main

import "fmt"

// Arrays is an introduction into to arrays in Golang.  It shows different ways of
// initializing an array and iterating over its contents.
func main() {
  // Array creation - 2
	var x [3]int = [3]int{4, 5, 6}
	fmt.Println("x is ", x)

  // Array creation - 1
	y := [3]int{1, 2, 3}
	fmt.Println("y is ", y)

  // Array creation - 3
	z := [3]int{7,
		8,
		9}
	fmt.Println("z is ", z)
  
  // Iterating over an array
  for i := 0; i < len(x); i++ {
    fmt.Println(x[i]*x[i])
  }
  
  // Iterating over an array - Take 2
  for _, value := range x {
    fmt.Println(value*value)
  }
}
