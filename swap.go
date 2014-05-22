package main
import "fmt"

// swapping two ints using pointers.

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	x := 13
	y := 3244
	fmt.Printf("Pre swap : x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Post swap: x = %d, y = %d\n", x, y)
}
