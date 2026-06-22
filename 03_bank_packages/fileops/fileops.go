package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func ReadBalance(balanceFile string) float64 {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000
	}
	readBalance := string(data)
	balance, _ := strconv.ParseFloat(readBalance, 64)

	return balance
}

func WriteBalance(balanceFile, balance string) float64 {
	os.WriteFile(balanceFile, []byte(balance), 0644)
	fmt.Printf("New balance: %v \n", balance)
	currentBalance, _ := strconv.ParseFloat(balance, 64)
	return currentBalance
}
