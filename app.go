package main // all files in this package will have access to these functions (i.e. main.go has access to these functions)

import (
	"fmt"
	"profit_calculator/fileops"
)

func app(balance *float64) {
	// infinite loop 
	for {
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your current balance is %.2f ZAR\n", *balance)
			fmt.Println()
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("You must deposit an amount more than 0!")
				fmt.Println()
				continue
			}
			*balance += deposit
			fmt.Printf("Your current balance is %.2f ZAR\n", *balance)
			fileops.WriteBalance(*balance)
			fmt.Println()
		} else if choice == 3 {
			fmt.Print("Your withdrawal: ")
			var withdrawal float64
			fmt.Scan(&withdrawal)
			if withdrawal > *balance || withdrawal < 0 {
				fmt.Println("You cannot withdraw more than you have or less than 0!")
				fmt.Println()
				continue
			}
			*balance -= withdrawal
			fmt.Printf("Your current balance is %.2f ZAR\n", *balance)
			fileops.WriteBalance(*balance)
			fmt.Println()
		} else if choice == 4 {
			profit()
			fmt.Println()
		} else {
			fmt.Println("Thank you for using Go bank")
			break
		}
	}
}
func profit() {
	var revenue, expenses, tax float64
	revenue = getInput("Enter your revenue amount in ZAR: ")
	expenses = getInput("Enter your expense amount in ZAR: ")
	tax = getInput("Enter your tax rate: ")

	EBT, profit, ratio := calcProfit(revenue, expenses, tax)
	fileops.WriteResults(EBT, profit, ratio)
	display(EBT, profit, ratio)
}

func getInput(inputText string) float64 {
	var input float64
	fmt.Print(inputText)
	fmt.Scan(&input)
	return input
}

func calcProfit(revenue, expenses, tax float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = revenue - (revenue * (tax / 100))
	ratio = EBT / profit
	return
}

func display(EBT, profit, ratio float64) {
	fmt.Print("Your Earnings Before Tax (EBT) is ")
	fmt.Println(EBT, "ZAR")
	fmt.Print("Your Profit is ")
	fmt.Println(profit, "ZAR")
	fmt.Printf("Your Profit Margin is %.3f\n", ratio)
}