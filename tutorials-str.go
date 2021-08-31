package main

import "fmt"

func //main() {
	var greeting = "hello World \n"
	fmt.Printf("normal string \n")
	fmt.Printf("%s", greeting)

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i]) //%x in output ascii value
	}
	fmt.Printf("\n")
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //quotes value
	fmt.Printf("%+q", sampleText)

	/*var n [10]int
	var i, j int //array
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element [%d]= %d\n", j, n[j])
	}
	*/

}
