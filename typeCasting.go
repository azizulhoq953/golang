// package main

// import "fmt"
// func main() {
// 	var sum int = 20
// 	var count int = 8
// 	var mean float32

// 	mean = float32(sum) / float32(count)

// 	fmt.Printf("Value of mean : %f\n", mean)
// }

package main

import (
	"fmt"
	"strconv"
)

//type casting TpointPractice
func main() {

	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "10.123"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newInt, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newInt)

	newfloat, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(newfloat)
}

// /* define an interface */
// type Shape interface {
// 	area() float64
// }

// /* define a circle */
// type Circle struct {
// 	x, y, radius float64
// }

// /* define a rectangle */
// type Rectangle struct {
// 	width, height float64
// }

// /* define a method for circle (implementation of Shape.area())*/
// func (circle Circle) area() float64 {
// 	return math.Pi * circle.radius * circle.radius
// }

// /* define a method for rectangle (implementation of Shape.area())*/
// func (rect Rectangle) area() float64 {
// 	return rect.width * rect.height
// }

// /* define a method for shape */
// func getArea(shape Shape) float64 {
// 	return shape.area()
// }

// func main() {
// 	circle := Circle{x: 0, y: 0, radius: 5}
// 	rectangle := Rectangle{width: 10, height: 5}

// 	fmt.Printf("Circle area: %f\n", getArea(circle))
// 	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
// }
