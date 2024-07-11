package main

import (
	"fmt"

	"globallogic.com/bank/customops"
)

const fileName = "balance.txt"

func main() {
	fmt.Println("Welcome to Go bank!")
	for {
		var accountBalance, err = customops.ReadFromFile(fileName)
		if err != nil {
			fmt.Println("Error!")
			fmt.Println(err)
			fmt.Println("------------")

		}

		communication() //This function is a part of communication.go file

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Balace is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid operation.")
				// return
				continue
			}
			accountBalance = accountBalance + depositAmount // accountBalance += depositAmount
			fmt.Println("Updated successfully! New balance: ", accountBalance)
			customops.WriteToFile(accountBalance, fileName)
		} else if choice == 3 {
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid operation.")
				return
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance.")
				return
			}
			accountBalance = accountBalance - withdrawAmount // accountBalance -= depositAmount
			fmt.Println("Withdraw successul! New balance: ", accountBalance)
			customops.WriteToFile(accountBalance, fileName)
		} else {
			// return
			break
		}
	}
	fmt.Println("Thanks for choosing Go banking!")
}
