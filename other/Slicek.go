package main

import "fmt"

func main() {
	company := []string{
		"Azizulhoq.co\t", "Branch:Bangladesh and America \t", "Japan \t",
	}

	company = append(company, "Append", "Append2") //append in last index
	company[2] = "Swithzerland"
	fmt.Println(company, "\n")
	fmt.Println(len(company), "\n")
	fmt.Println(cap(company))
}
