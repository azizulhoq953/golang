package main

import "fmt"

func main() {
	company := []string{
		"Azizulhoq.co\t", "Branch:Bangladesh and America \t", "Japan \t",
	}

	company_Earn := []string{
		"55680", "78000", "99000",
	}

	company_all := append(company, company_Earn...)

	company = append(company, "Append", "Append2") //append in last index
	company[2] = "Swithzerland"
	fmt.Println(company_all, "\n")
	fmt.Println(len(company_all), "\n")
	fmt.Println(cap(company_all))
}
