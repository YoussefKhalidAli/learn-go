package main

import (
	"fmt"
	"math"
)

func main() {
	inflation := 4.5
	var investment, years, expectedReturn float64

	fmt.Print("Enter investment: ")
	fmt.Scan(&investment)

	fmt.Print("Enter expectedReturn: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	returned := investment * math.Pow(1+expectedReturn/100, years)
	realReturned := returned / math.Pow(1+inflation/100, years)
	fmt.Println(returned)
	fmt.Println(realReturned)
}
