package main

import "fmt"

func //main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("address stored in of variable: %d\n", &a)

	fmt.Printf("address stored in ip variable : %x\n", *ip)

	fmt.Printf("Value of pointer access *ip variable: %d\n ", *ip)
}
