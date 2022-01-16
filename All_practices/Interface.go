//interface

package main

import "fmt"

type Email interface{
write(string,string,string,string)
send()
read()

}

type Test struct{
To string
From string
Subject string
Message string
}

func (t Test)write(to string,from string,subject string,Message string){

fmt.Println(t.To ,from, subject,Message)

}

func (t Test)send(){

fmt.Println(t.To, "Email Sent")

}



func (t Test) read(){

fmt.Println(t.From,"Recived")

}

func main(){

//var e  Email

//e.write"Candidte Send"
//fmt.Println(e)

var tst Test

tst.To="azizulhoq4305@gmail.com\n"

tst.From="azizurrahman4305@gmail.com\n"

tst.Subject="test Email\n"

tst.Message="Hello This Is My InterFace First Project test Email\n"

tst.write (tst.To, tst.From,tst.Subject,tst.Message)

}