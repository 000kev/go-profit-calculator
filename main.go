package main

import "fmt"

func main() {
	var revenue, expenses, tax float64
	fmt.Print("Enter your revenue amount in ZAR: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expense amount in ZAR: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&tax)

	EBT := revenue - expenses
	profit := revenue - ( revenue * (tax/100))
	ratio := EBT / profit

	fmt.Print("Your Earnings Before Tax (EBT) is ")
	fmt.Println(EBT, "ZAR")
	fmt.Print("Your Profit is ")
	fmt.Println(profit, "ZAR")
	fmt.Print("Your Ratio is ")
	fmt.Println(ratio)
}