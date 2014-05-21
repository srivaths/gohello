package main
import "fmt"
func main() {
	// A map where the keys are strings and the associated values are ints
	var x = make(map[string] int)
	x["sri"] = 50
	x["jimmy"] = 3
	fmt.Println(x)

	// Another way to initialize a map
	y := map[string] int {
		"bubba": 93,
		"trickster": -1,
	}
	fmt.Println("y is ", y)

	// Whoa!
	z := map[string] map[int] string {
		"fruits": map[int] string {
			50: "apple",
			90: "banana",
		},
		"cars": map[int] string {
			11: "ford",
			19: "caddy",
			39: "lancia",
		},
	}
	fmt.Println(z)

}
