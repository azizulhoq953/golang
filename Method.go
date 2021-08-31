//paractice

package main

import "fmt"

/*type AirLines struct{

pilot string

AirHostess string

setNumber string

Passenger string

Ticket string

Pick string


}


func (A AirLines)getpilot()string{

return  A.pilot

}

func main(){

var A AirLines

A.pilot="Azizul Hoq"
A.AirHostess="First Level:Tanha,Sceond Level:Reyana..."
A.setNumber="Vip:1,Vip:2...normal: 1,2,3,4,5,6,7,8,9,10"

A.Passenger="Vip:1.Atiqur Rahman,Vip:2.Sadiqa.."
A.Pick="USA"

fmt.Println( A,A.pilot)
}*/
//InterFace

type AirLines interface{

Airport (string,string,string,string )

Biman (string,string,string)

}

type Airport struct{

AirForce string

AirOrganization string

salary string

}

type Biman struct{

pilot string

AirHostess string

passenger string


}


func (A Airport)getAirOrganization()string{

return A.AirOrganization


}


func main(){

var A Airport

A.AirForce="\n\nAirForce:Is Duity Morning And Evening And Night 120 This Unit\n\n\n"

A.AirOrganization="AirLines:USA, Emirates, British,Canada,Bangladesh,china\n\n\n"

A.salary="Salary:1stLevel:1200$,2ndLevel:800$,3rdLevel:600$\n\n\n"

var B Biman 

B.pilot="AzizulHoq,Galib\n\n"

B.AirHostess="Jebunnesa,Jaren Afroj,Tasnia,Tabib..\n\n"

B.passenger="People:111\n\n"

fmt.Println(A,A.AirOrganization)
fmt.Println(B)
}




