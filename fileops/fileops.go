package fileops // to create new packages, you have to put them in their own subfolder
// this package is then imported by other files 

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

// in Go, functions that are exported must start with an Upercase letter
func ReadBalance() (float64, error) {
	data, err := os.ReadFile("balance.txt") // underscore tells Go you're not using the other variable, namely error in this case
	if err != nil {
		return 1000, errors.New("failed to find balance file") // error handling
	}

	bal := string(data)
	balance, err := strconv.ParseFloat(bal, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value") // error handling
	}
	return balance, nil
}


func WriteBalance(balance float64) {
	bal := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(bal), 0644) // 0644 is a way of defining read and write permission
}

func WriteResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.3f\nProfit: %.3f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}