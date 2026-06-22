package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	print("Enter revenue: ")
	fmt.Scan(&revenue)

	print("Enter expenses: ")
	fmt.Scan(&expenses)

	print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax / taxRate)

	fmt.Printf("Earnings before tax: %v \n", earningsBeforeTax)

	fmt.Printf("Type of Earnings after tax:  %T \n", earningsAfterTax)
	fmt.Printf("Earnings after tax:  %#v \n", earningsAfterTax)

	fmt.Printf("ratio: %.2f", earningsBeforeTax/earningsAfterTax)
}
