package main

import "fmt"

func main() {
	var name int

Loop:
	fmt.Printf("Your age Is Not restrict:\n")
	fmt.Printf("Please Contact For Support:")

	fmt.Printf("Enter Your Age:")
	fmt.Scanln(&name)
	if name <= 19 {
		goto Loop
	} else {
		fmt.Printf("Your age is Comfort:")
	}
}

// package main

// import "fmt"

// func main() {
// 	var input int
// Loop:
// 	{
// 		fmt.Print("You Are Not Eligible to Vote:\n")
// 	}
// 	fmt.Print("Enter Your Age:\n")
// 	fmt.Scanln(&input)

// 	if input <= 17 {
// 		goto Loop

// 	} else {
// 		fmt.Print("You Can Vote:")
// 	}
// }
