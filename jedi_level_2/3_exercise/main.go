package main
// Typed and Untyped Constant
import "fmt"

const (
	a = 4
	b = 8
	c float64 = 43.0
)

func main() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}