package main
import ("fmt"; "math")

// Shape interface defines properties of a shape.
type Shape interface {
  Permimeter() float64
  Area() float64
}

// Rectangle data type.
type Rectangle struct {
  // x,y: coordinates of bottom left
  // l,w: Length & width
  x, y, l, w float64
}

// Area method for a Rectangle
func (r *Rectangle) Area() float64 {
  return (r.l * r.w)
}

// Perimeter method for a Rectangle
func (r *Rectangle) Perimeter() float64 {
  return (4 * r.l * r.w)
}

type Circle struct {
  x, y, r float64
}

// Area method for a Circle
func (c Circle) Area() float64 {
  return math.Pi * c.r * c.r
}

// Perimeter method for a Circle
func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.r
}

func main() {
  c := Circle{0, 0, 5}
  r := Rectangle{0, 0, 4, 8}
  fmt.Printf("Circle:\n\tRadius: %f\n\tArea     : %f\n\tPerimeter: %f\n", c.r, c.Area(), c.Perimeter())
  fmt.Printf("Rectangle:\n\tLength, width: %f, %f\n\tArea: %f\n\tPerimeter: %f\n", r.l, r.w, r.Area(), r.Perimeter())
}
