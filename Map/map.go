package main

import "fmt"

func main() {
	//# 1
	var countryCapitalMap map[string]string

	//#2
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["CN"] = "北京"
	countryCapitalMap["JP"] = "东京"
	countryCapitalMap["US"] = "华盛顿"

	for country := range countryCapitalMap {
		fmt.Println(country, "caption is: ", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["US"]
	if (ok) {
		fmt.Println("OK", capital)
	} else {
		fmt.Println("No")
	}

	delete(countryCapitalMap, "JP")
	fmt.Println("del jp")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

}
