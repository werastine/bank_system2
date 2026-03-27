package main

import bank "go_coding/bank_system2/Bank"

func main() {
	MyAccount := bank.NewBankAccount()
	MyAccount.CreateCard("Arsenii Zhadanenko", "GP99011")
}

//Create Card num creation,
//Finish profile setter, it will storage name, and user's data
// Start thinking about payments
