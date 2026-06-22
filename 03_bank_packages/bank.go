package main

import (
	"fmt"

	"learn.com/bank-packages/fileops"
)

var balanceFile = "balance.txt"
var currentBalance = fileops.ReadBalance(balanceFile)

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
			fmt.Printf("Your balance is %v \n", fileops.ReadBalance(balanceFile))
		case 2:
			var depositAmount float64
			print("Deposit amount: ")
			fmt.Scan(&depositAmount)
			deposit := fmt.Sprint(currentBalance + depositAmount)
			currentBalance = fileops.WriteBalance(balanceFile, deposit)
		case 3:
			var withdrawlAmount float64
			print("Withdrawl amount: ")
			fmt.Scan(&withdrawlAmount)
			deposit := fmt.Sprint(currentBalance - withdrawlAmount)
			currentBalance = fileops.WriteBalance(balanceFile, deposit)
		case 4:
			println("Thank you for choosing moai bank")
			return
		}
	}
}
