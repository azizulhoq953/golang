package main

import "fmt"

func main() {
	number := []int{5, 4, 5, 7, 3, 2, 1}
	for i := range number {
		fmt.Println("Slice", i, "is", number[i])
	}

	//map
	countryMap := map[string]string{"bangladesh": "Dhaka", "China": "tokyo", "India": "Mumbai", "America": "lossAngel"}

	//print Map using key
	for country := range countryMap {
		fmt.Println("capital of", country, "is", countryMap[country])

	}

}
