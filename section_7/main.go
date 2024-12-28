package main

import "fmt"

type floatMap map[string]float64

func main() {
	// lists()
	// maps()

	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Henry")

	fmt.Println(userNames)

	coursesRatings := make(floatMap, 3)

	coursesRatings["go"] = 4.7
	coursesRatings["react"] = 4.8
	coursesRatings["angular"] = 4.7

	fmt.Println(coursesRatings)

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range coursesRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
