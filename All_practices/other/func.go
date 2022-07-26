//using struct
package main

import "fmt"

type helipad struct {
	name     string
	area     string
	location string
	company  string
	budget   int
}

func main() {
	var hp1 helipad
	var hp2 helipad
	var hp3 helipad
	hp1.name = "Bangladesh AirPort"
	hp1.area = "Cox's bazer"
	hp1.location = "Bangladesh"
	hp1.company = "AirLiness.co"
	hp1.budget = 89900000

	/*hp2 value Decleared*/

	hp2.name = "Bangladesh AirPort"
	hp2.area = "Barishal Sadar"
	hp2.location = "Bangladesh"
	hp2.company = "AirLiness.co"
	hp2.budget = 80000000

	hp3.name = "Bangladesh AirPort"
	hp3.area = "Kuakata"
	hp3.location = "Bangladesh"
	hp3.company = "AirLiness.co"
	hp3.budget = 89000000
	//print hp
	printhelipad(&hp1)

	//print hp2
	printhelipad(&hp2)

	printhelipad(&hp3)

	//fmt.Println(hp)
}

func printhelipad(helipads *helipad) {
	fmt.Printf("Helipad Name:%s\n", helipads.name)
	fmt.Printf("Helipad area:%s\n", helipads.area)
	fmt.Printf("Helipad location:%s\n", helipads.location)
	fmt.Printf("Helipad Sponsored company:%s\n", helipads.company)
	fmt.Printf("Helipad Budget:%d\n", helipads.budget)
	fmt.Printf("\n")
}
