package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := input("revenue")
	handleError(err)
	expenses, err := input("expenses")
	handleError(err)
	taxRate, err := input("tax rate")
	handleError(err)

	earningsBeforeTax, earningsAfterTax, ratio := earnings(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %v \n", earningsBeforeTax)

	fmt.Printf("Type of Earnings after tax:  %T \n", earningsAfterTax)
	fmt.Printf("Earnings after tax:  %#v \n", earningsAfterTax)

	fmt.Printf("ratio: %.2f", ratio)
}

func input(entry string) (float64, error) {
	var entryValue float64
	fmt.Printf("Enter %v:", entry)
	fmt.Scan(&entryValue)

	if entryValue <= 0 {
		return 0, errors.New("Invalid inuput. Must be greater than 0")
	}
	return entryValue, nil
}

func earnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt - (ebt / taxRate)
	return ebt, eat, ebt / eat
}

func handleError(err error) {
	println(err.Error())
	os.Exit(1)
}
