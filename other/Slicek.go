package main

import "fmt"

func main() {
	company := []string{
		"Azizulhoq.co", "Branch:Bangladesh and America",
	}

	fmt.Println(company, "\n")
	fmt.Println(len(company), "\n")
	fmt.Println(cap(company))
}
