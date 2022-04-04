/* explaination to problem
Make a function that will return a greeting statement that uses an input; your program should return, "Hello, <name> how are you doing today?".

[Make sure you type the exact thing I wrote or the program may not execute properly]



*/

// package kata

// func Greet(name string) string {
// 	// fill in solution here
// 	return "Greet"
// }

// func main() {

// 	Greet("Hello")

// }

/*
palindrome

*/

package main

func isPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}
