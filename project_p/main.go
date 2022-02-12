package main

import("fmt")

func main(){
/*
    //int data type

var n int=10;
var m int=20;
var sum int
sum=n+m;
fmt.Println(sum);
*/

	//string data type
var Firstname string;

var SecondName string;

var name string;

fmt.Scanln(&Firstname); //usear input

fmt.Scanln(&SecondName) //usear input

name=Firstname+SecondName;

fmt.Printf("%v ",name) //print data

}