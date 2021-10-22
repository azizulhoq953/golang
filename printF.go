/*The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

%v is used to print the value of the arguments
%T is used to print the type of the arguments*/
package main

import (
	"fmt"
)

func main() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v\n i is type: %T\n", i, i)
	fmt.Printf("j has value: %v%%\n j is type: %T\n", j, j)
	fmt.Printf(" i is syntax %#v\n", i)
	fmt.Printf(" j is syntax %#v\n", j)
}

/*The Println() function is similar to Print() with the difference that a whitespace is added between the arguments,
and a newline is added at the end:*/
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i, j)
// }

/*Print() inserts a space between the arguments if neither are strings:*/

// package main
// import ("fmt")

// func main() {
//   var i,j = 10,20

//   fmt.Print(i,j)
// }
