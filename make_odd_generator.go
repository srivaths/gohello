package main
import "fmt"
// http://www.golang-book.com/7/index.htm
// Generates odd numbers.
// Things of note:
//  - Returns a function
//  - The value of i is persistent
func makeOddGenerator() func() uint {
  var i uint = 1
  return func() (retval uint) {
    retval = i
    i += 2
    return
  }
}

func main() {
  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
}