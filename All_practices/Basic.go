package main

import "fmt"

func main() {
	var input int

	fmt.Print("Enter Your Value 1 to 100:")
	fmt.Scanln(&input)
	if input < 0 || input > 100 {
		fmt.Print("Please Enter Valid Number:")
	} else if input >= 0 && input <= 33 {
		fmt.Print("Your Grde Of:D\n")
		fmt.Print("Your GPA :1", "\n")
	} else if input >= 33 && input <= 40 {
		fmt.Printf("Your Grade Of: C \n")
		fmt.Printf("Your GPA : 2")
	} else if input > 40 && input <= 50 {
		fmt.Printf("Your Grade Of: B\n")
		fmt.Printf("Your GPA: 3")
	} else if input >= 50 && input <= 60 {
		fmt.Print("Your Grade Of: A\n")
		fmt.Print("Your GPA:4.00")
	} else {
		fmt.Print("GPA:5")
	}
	return 0
}
