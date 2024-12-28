package main

import "fmt"

func lists() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)

	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)
}

// func lists() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)

// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
