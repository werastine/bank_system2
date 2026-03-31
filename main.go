package main

import (
	"fmt"
	bank "go_coding/bank_system2/Bank"
)

func main() {
	MyAccount := bank.NewBankAccount("Arsenii Zhadanenko", 17, "GP99011")
	MyAccount.CreateCard()
	fmt.Println(MyAccount.UserCard)
	fmt.Println(MyAccount.UserProfile.UserID)
	MyAccount.CreateCard()

}

//Finish profile setter, it will storage name, and user's data
// Start thinking about payments

//Move id storage to bank account, cause now id storage is in profile data, что не корректно
