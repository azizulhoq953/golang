package main

import"fmt"

	func update(d *int, e *string){
	*d=*d+5
	*e=*e+"Hoq"
	fmt.Println(e)
	//fmt.Println(d)
	
	return
}

func main(){
	
//var a *int
	//var j int=10
	
	//a=&j
	
	f:=20
	g:="Azizul"
	
	
	
	update(&f,&g)
	
	//fmt.Println(&j,*a)
	fmt.Println(f,g,&g)


}