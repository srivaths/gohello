package main
import "fmt"

// Find the smallest int in the given array
func main() {
  x := []int{
      48,96,86,68,
      57,82,63,70,
      37,34,83,27,
      19,97, 9,17,
  }
  var smallest = x[0]

  for _, value := range x {
    if value < smallest {
      smallest = value
    }
  }
  fmt.Printf("The smallest value in %v is %d\n", x, smallest)
}