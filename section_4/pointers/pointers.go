package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", agePointer)
	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	*age = *age - 18
	return *age
}
