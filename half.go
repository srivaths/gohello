package main
import "fmt"

// Determines half of given value and indicates whether given value is even.
func half(x int) (int, bool) {
  y := x/2
  even := false
  if x % 2 == 0 {
    even = true
  }
  return y, even
}

func main() {
  x := 7
  halfvalue, isEven := half(x)
  fmt.Printf("Is %d even? Answer: %v.  Half of %d is %d\n", x, isEven, x, halfvalue)
}