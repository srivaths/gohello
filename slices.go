package main
import "fmt"

// A slice is a segment of an array. Like arrays slices are indexable and have
// a length. Unlike arrays this length is allowed to change. 
func main() {
  // A empty zero-size slice
  var x[]int
  fmt.Println(x) // just to please the compiler
  
  // The preferred way to create a slice is using the built-in make function
  var y = make([]int, 8, 20)
  // This created a slice of size 8 ints with the capacity to hold 20 ints
  // TODO -- The third parameter is confusing.  Since the length of a
  // slice is allowed to change why specify it?
  fmt.Println(y) // just to please the compiler
  
  var z = [3]int { 3, 4, 5,}
  
  fmt.Println("z is a slice[3] = ", z)
  
  // And just for grins see how the compiler ignores the extra comma
  // at the end!
  
  // Another way to declare a slice is to give the start & end
  // indices of a source array
  
  source := []int { 90, 91, 92, 93, 94, 95, 18, 11, 33}
  dest := source[2: 6]
  fmt.Println("source        = ", source)
  fmt.Println("source[2:6]   = ", dest)
  
}