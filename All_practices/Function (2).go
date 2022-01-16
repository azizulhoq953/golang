package main

import "fmt"

func rectangle(l,b int) (area, parameter int){

parameter=2*(l+b)
area=l*b

return


}

func main(){

	a,p:=rectangle(10,20)
	
	fmt.Println(a,p)


}