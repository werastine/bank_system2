package main

import (
	"fmt"
	bankpkg "go_coding/bank_system2/Bank"
)

func main() {
	bank := bankpkg.NewBank()
	MyAccount := bankpkg.NewBankAccount(bank, "Arsenii Zhadanenko", 17, "GP99011")
	MyAccount.CreateCard()
	fmt.Println(MyAccount.UserCard)
	fmt.Println(MyAccount.UserProfile.UserID)
	MyAccount.CreateCard()

}

// Del cards & id's storage from bank accaunt and implement them into Bank struct
// so additional task is to create bank structure and prepare code

//Finish profile setter, it will storage name, and user's data
// Start thinking about payments

//Move id storage to bank account, cause now id storage is in profile data, что не корректно
