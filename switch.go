// Used for Switch Case
package main

import "fmt"

func main(){

fmt.Print("Enter Your Age=")

var age int

fmt.Scanf("%d",&age)

switch age {
			case 1,2,3:
			fmt.Println("Infant")
			fallthrough //common down case Print
			
			
	
			case 4,5,6,7,8,9,10,12,13,14:
			fmt.Println("Children")
			
			case 15,16,17,18,19,20:
			fmt.Println("Teen age")
			
			case 21,22,23,24,25,26,27,28,29,30:
			fmt.Println("Adult")

			default:
			fmt.Println("Adult")
}
}