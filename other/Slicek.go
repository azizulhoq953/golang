package main

import "fmt"

func main() {
	company := [10]string{
		"Azizulhoq.co", "Branch:Bangladesh and America", "Japan",
	}

	company[3] = "Swithzerland"
	fmt.Println(company, "\n")
	fmt.Println(len(company), "\n")
	fmt.Println(cap(company))
}
