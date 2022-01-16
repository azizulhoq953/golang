package main

import (
	"fmt"
	"reflect"
)

func //main() {

	//string Array

	//student := [...]string{"Azizul", "Bangladesh"}
	//fmt.Println(student, len(student))

	//Slice
	//var students [5]string

	//students[0] = "Azizul"
	//students[1] = "Anonya"
	//students[2] = "Asgor"
	//students[3] = "Mostain"
	//x := students[0:3]
	//fmt.Println(x, len(x), len(students))

	//y := make([]string,6)
	//fmt.Println(y)

	//var rose []string

	//rose = append(rose, "jui", "Belly", "joba")

	//fmt.Println(rose)
	//fmt.Printf("%T\n", rose)
	//fmt.Printf("%T", students)

	//Sceond Practice

	var Country []string

	Country = append(Country, "America\n", "Canada\n", "Switzerland\n", "Finland\n", "ArabEmarates\n")

	//fmt.Println(Country, len(Country))

	//fmt.Printf("%T", Country)

	b := reflect.TypeOf(Country).Kind().String()
	fmt.Println(b)

}
