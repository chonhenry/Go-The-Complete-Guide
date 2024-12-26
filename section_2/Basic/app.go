package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	furtureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// furtureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := furtureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf("Furture Value: %.1f\n", furtureValue)
	formattedRFV := fmt.Sprintf("Furture Value (adjusted for inflation): %.1f\n", futureRealValue)

	// fmt.Println("Furture Value: ", furtureValue)
	// fmt.Println("Furture Value (adjusted for inflation): ",futureRealValue)
	// fmt.Printf("Furture Value: %.1f\nFurture Value (adjusted for inflation): %.1f\n", furtureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}