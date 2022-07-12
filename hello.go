/* Package declaration*/
package main
/* preprocessor*/
import "fmt"
import "math"
func main() {
	var a, b, c = 1, 1, "sarvesh"
	var x float64 = 20.0
	y := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	l, k := two("hello", "sarvesh")
	fmt.Println(l, k)
	fmt.Println(split(17))
	forLoop()
	whileInGoIsFor()
	sqrt(4)
}
// Multiple Return Function
func two(greet, name string) (string, string) {
	return greet, name
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum of 10 number is %v\n", sum)
}
func whileInGoIsFor() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}