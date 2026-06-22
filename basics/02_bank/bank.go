package main

import (
	"fmt"
	"os"
	"strconv"
)

var balanceFile = "balance.txt"
var currentBalance = readBalance()

func main() {
	println("Welcome to moai bank")

	var choice int

	for {
		println("What brings you here?")
		println("1. Check balance")
		println("2. Make a deposit")
		println("3. Make a withdrawk")
		println("4. Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %v \n", readBalance())
		case 2:
			var depositAmount float64
			print("Deposit amount: ")
			fmt.Scan(&depositAmount)
			deposit := fmt.Sprint(currentBalance + depositAmount)
			writeBalance(deposit)
		case 3:
			var withdrawlAmount float64
			print("Withdrawl amount: ")
			fmt.Scan(&withdrawlAmount)
			deposit := fmt.Sprint(currentBalance - withdrawlAmount)
			writeBalance(deposit)
		case 4:
			println("Thank you for choosing moai bank")
			return
		}
	}
}

func readBalance() float64 {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000
	}
	readBalance := string(data)
	balance, err := strconv.ParseFloat(readBalance, 64)

	if err != nil {
		panic(err)
	}

	return balance
}

func writeBalance(balance string) {
	os.WriteFile(balanceFile, []byte(balance), 0644)
	fmt.Printf("New balance: %v \n", balance)
	currentBalance, _ = strconv.ParseFloat(balance, 64)
}
