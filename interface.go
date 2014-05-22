package main
import ("fmt"; "math")
type Shape interface {
  Permimeter() float64
  Area() float64
}

type Rectangle struct {
  // x,y: coordinates of bottom left
  // l,w: Length & width
  x, y, l, w float64
}
func (r *Rectangle) Area() float64 {
  return (r.l * r.w)
}
func (r *Rectangle) Perimeter() float64 {
  return (4 * r.l * r.w)
}

type Circle struct {
  x, y, r float64
}
func (c Circle) Area() float64 {
  return math.Pi * c.r * c.r
}
func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.r
}

func main() {
  c := Circle{0, 0, 5}
  r := Rectangle{0, 0, 4, 8}
  fmt.Printf("Circle:\n\tRadius: %f\n\tArea     : %f\n\tPerimeter: %f\n", c.r, c.Area(), c.Perimeter())
  fmt.Printf("Rectangle:\n\tLength, width: %f, %f\n\tArea: %f\n\tPerimeter: %f\n", r.l, r.w, r.Area(), r.Perimeter())
}