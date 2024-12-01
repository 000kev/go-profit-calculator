package main

import "fmt"

func main() {
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
	profit = revenue - ( revenue * (tax/100))
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