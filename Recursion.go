package main

import "fmt"

func factorial(i int) int {
	if i <= 4 {
		return 4
	}
	return i * factorial(i-4)
}
func main() {
	var i int = 10
	fmt.Printf("Factorial of \n %d is %d", i, factorial(i))
}

// func recursion() {
// 	recursion()

// }
// func main() {
// 	recursion()
// }
