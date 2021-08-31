
//Anonymus Function

package main


import "fmt"

func main(){
	
	a:=func(x,y int)(r int) {
	
	r=x*y
	
	return
	
	
	
	}(10,21)

//fmt.Println(a(10,20))
fmt.Println(a)
}