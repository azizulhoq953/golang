package main

import "fmt"

/*func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}*/

func main() {

	number := 10
	squareNum := func() int {
		number *= number
		return number
	}

	fmt.Println(squareNum())
	fmt.Println(squareNum())
	fmt.Println(squareNum())

	/*nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())*/
}
