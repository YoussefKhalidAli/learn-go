package main

import (
	"fmt"
)

func main() {

	revenue := input("revenue")
	expenses := input("expenses")
	taxRate := input("tax rate")

	earningsBeforeTax, earningsAfterTax, ratio := earnings(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %v \n", earningsBeforeTax)

	fmt.Printf("Type of Earnings after tax:  %T \n", earningsAfterTax)
	fmt.Printf("Earnings after tax:  %#v \n", earningsAfterTax)

	fmt.Printf("ratio: %.2f", ratio)
}

func input(entry string) float64 {
	var entryValue float64
	fmt.Printf("Enter %v:", entry)
	fmt.Scan(&entryValue)
	return entryValue
}

func earnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt - (ebt / taxRate)
	return ebt, eat, ebt / eat
}
