package main

import "fmt"


/*type Phone struct{

Apps string

Ram string

Rom string
Price string
Warranty string

}*/

func main(){
/*var p1 Phone
p1.Apps=" Apps:Phone History"
p1.Ram="4GB"
p1.Rom="64GB Rom"
p1.Price="32000 kip Price"
p1.Warranty="1 Years Replacemant Gurranty 2 Years Warranty"

fmt.Println(p1)
fmt.Println(p1.Apps)
fmt.Println(p1.Ram)
fmt.Println(p1.Rom)
fmt.Println(p1.Price)
fmt.Println(p1.Warranty)
*/
//anonymus struct
b1:=struct{

Apps string

Ram string

Rom string
Price string
Warranty string
}{
Apps: "Apps Phone History",

Ram :"4GB Ram",
Rom :"64GB Rom",
Price: "32000 kip Price",
Warranty: "1 Years Replacemant Gurranty 2 Years Warranty",


}

fmt.Println(b1)
fmt.Println(b1.Apps)
fmt.Println(b1.Ram)
fmt.Println(b1.Rom)
fmt.Println(b1.Price)
fmt.Println(b1.Warranty)


}