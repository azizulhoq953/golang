package main

import "fmt"

// func myFunction(x, y string) string {
// 	return x + y
// }

// func main() {
// 	fmt.Println(myFunction("azizul", "Hoq"))
// }

func Number(x, y, z float64) float64 {
	return x / y / z
}

func main() {
	fmt.Println("Division Of:", Number(10, 3, 1))
}

//Net Topics Start Recursion
