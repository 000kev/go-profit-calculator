package main

import "fmt"

func main() {
	var balance float64 = 1000

	fmt.Println("Welcome to Go Bank!\nWhat woutld you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Check profit")
	fmt.Println("5. Exit")

	// Go only has for loops and not while like other languages
	// for i := 0; i <200; i++ {}  => finite loop
	// infinite loop =>
	app(&balance)
	fmt.Println("We hope to see you again!")
}

func app(balance *float64) {
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
