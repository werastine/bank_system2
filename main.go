package main

import (
	"fmt"
	bank "go_coding/bank_system2/Bank"
)

func main() {
	MyAccount := bank.NewBankAccount()
	MyAccount.CreateCard("Arsenii Zhadanenko", "GP99011")
	fmt.Println(MyAccount.UserCard)
	MyAccount.CreateCard("Arsenii Zhadanenko", "GP99011")

}

//Create duplicate detector, map will contain user id and card number
//Finish profile setter, it will storage name, and user's data
// Start thinking about payments
