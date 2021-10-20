package main

import (
	"fmt"
)

func address(address int, Name string) {
	fmt.Println(address, Name)
}

func myName(fname string) {
	fmt.Println(fname)
}

func main() {
	myName("Name:Azizul")

	address(22, "age")
}
