package main

import "fmt"

func main() {

	var countryCapitalMap map[string]string
	/*Map Create*/

	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	countryCapitalMap["Bangladesh"] = "Dhaka"
	countryCapitalMap["italy"] = "Rome"
	countryCapitalMap["swithzerland"] = "Zurich"
	countryCapitalMap["japan"] = "tokyo"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	/* test if entry is present in the map or not */

	capital, ok := countryCapitalMap["Zurich"]

	/* if ok is true entry is present otherwise entry is absent*/

	if ok {
		fmt.Println("Capital of Zurich is", capital)
	} else {
		fmt.Println("capital Of Zurich is not Present:")
	}

}
