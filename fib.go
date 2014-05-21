package main
import "fmt"

// Implementatoin of Fibonacci numbers as an example of 
// recursion.


func fib(i int) int {
  if i == 0 {
 		return i
	}
 return i + fib(i-1)
}

func main() {
	fmt.Println(fib(10))
}
