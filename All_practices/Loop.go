//Used of for,range, loop

package main

import "fmt"

func main(){

//for loop
/*for i:=0; i<10;i++{
			fmt.Println(i)

}*/
//range loop
students :=[]string{ "Azizul Hoq","Asgor Ali","Anoyonna"}

for i, std:= range students {

	fmt.Println(i,std)
	}
}