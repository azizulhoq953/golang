package main

import "fmt"

type car struct {
	name    string
	driver  string
	model   string
	speed   float32
	set     int
	license string
}

func //main() {

	var cars1 car
	//multiple called  type
	var cars2 car

	//car 1 spwcification
	cars1.name = "BMW"
	cars1.driver = "azizul"
	cars1.model = "A305"
	cars1.speed = 3.500
	cars1.set = 4
	cars1.license = "21344nh3l"

	//cars2 specification
	cars2.name = "Mercedes"
	cars2.driver = "azizul"
	cars2.model = "Ab3056"
	cars2.speed = 3.555
	cars2.set = 3
	cars2.license = "21344nh65"

	fmt.Printf("My Carse name: %s\n", cars1.name)
	fmt.Printf("My Carse driver: %s\n", cars1.driver)
	fmt.Printf("My Carse Model: %s\n", cars1.model)
	fmt.Printf("My Carse speed: %f\n", cars1.speed)
	fmt.Printf("My Carse set: %d\n", cars1.set)
	fmt.Printf("My Carse license: %s\n", cars1.license)

	// second cars2
	fmt.Printf("My second carse \n")

	fmt.Printf("My Carse name: %s\n", cars2.name)
	fmt.Printf("My Carse driver: %s\n", cars2.driver)
	fmt.Printf("My Carse model: %s\n", cars2.model)
	fmt.Printf("My Carse speed: %f\n", cars2.speed)
	fmt.Printf("My Carse set: %d\n", cars2.set)
	fmt.Printf("My Carse license %s\n", cars2.license)
}
