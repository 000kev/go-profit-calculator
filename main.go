package main

import (
	"fmt"
	"profit_calculator/fileops" // to import your package, you need to include the path of your module i.e. profit-calculator
)

func main() {
	var balance, err = fileops.ReadBalance() // errors don't crash the app in Go but rather return empty byte strings by default '0'

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------")
		fmt.Println()
		// you could also terminate execution by using return or using the panic() function
	}
	fmt.Println("Welcome to Go Bank!\nWhat woutld you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Check profit")
	fmt.Println("5. Exit")

	// Go only has for loops and not while like other languages
	// for i := 0; i <200; i++ {}  => finite loop
	// there are also conditional loops => for choice != 5 {}
	// alternatively you can use switch statements
	// break has a special meaning in Go whereby it breaks out of the entire switch statement
	app(&balance)
	fmt.Println("We hope to see you again!")
}

