
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
