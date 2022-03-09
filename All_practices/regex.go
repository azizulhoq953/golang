//package main

// func main() {
// 	//string type of regex
// 	re := regexp.MustCompile(".com")
// 	// fmt.Println(re.FindString("google.com"))
// 	// fmt.Println(re.FindString("abc.org"))
// 	// fmt.Println(re.FindString("fb.com"))
// 	fmt.Println(re.FindString("google.com"))
// 	fmt.Println(re.FindStringIndex("youtube.com"))
// 	fmt.Println(re.FindStringSubmatch("pintarest.com"))

// 	fmt.Printf("\n\n")

// 	//int type of regex
// 	sampleRegexp := regexp.MustCompile(`\d`)
// 	fmt.Println("For regex \\d")
// 	match := sampleRegexp.MatchString("1")
// 	fmt.Printf("For 1: %t\n", match)

// }
package main

import (
	"fmt"
	"regexp"
)

func main() {
	string := "Ra13, Pa67, S-60 AC-21, UP44"
	expression := regexp.MustCompile(`[-]?\d[\d]*[\]?[\d{2}]*`)
	res := expression.FindAllString(string, 4)
	for _, attribute := range res {
		fmt.Println("The number with string is:", attribute)
	}
}
