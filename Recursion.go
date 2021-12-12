//Factorial
package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	var i int;
	fmt.Println("Enter Any Number:");
	fmt.Scanf("%d",&i);
	fmt.Printf("Factorial of %d is %d", i, factorial(i))
}

/* package main

import "fmt"

func fibonaci(i int) (ret int) { //fibonacci
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
} */

// package main

// import "fmt"

// func factorial(i int) int {
// 	if i <= 4 {
// 		return 4
// 	}
// 	return i * factorial(i-4) //factorial
// }
// func main() {
// 	var i int = 10
// 	fmt.Printf("Factorial of \n %d is %d", i, factorial(i))
// }

// func recursion() {
// 	recursion()

// }
// func main() {
// 	recursion()
// }
