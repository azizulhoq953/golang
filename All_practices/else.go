//Used of  if,else if,else

package main

import "fmt"

func main(){

fmt.Print("Enter Your Age=")

var age int

fmt.Scanf("%d",&age)
//fmt.Println(age)

if age < 3 {
	
fmt.Println("Baby")
	
}else if age>3 && age<12 {

		fmt.Println("Children")

}else if age>12 && age<19{

	fmt.Println("Teen age")

}else if 19<age && 40>age {
	fmt.Println("Adult")
}else{

	fmt.Println("Old Man")

}


}