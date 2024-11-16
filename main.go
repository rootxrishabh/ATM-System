package main

import (
	"fmt"
	"math/rand"
)

/*
The goal is to create a ATM system, here are the requirements -
	1. Authenticate via PIN
	2. Withdraw, deposit and transfer money
	3. Check account balance
	4. Support multiple accounts
	5. Cash despensing and error handling for insufficient funds.
*/

type Account struct {
	accNo    int
	availBal int
	accType  int
	pinNo    int
}

var (
	Accounts = make(map[int]Account)
)

const (
	savingsAccount = iota
	currentAccount
)

func Transfer(accNo, pin, transferAccNo, amount int) bool {
	if !Authenticate(accNo, pin) {
		fmt.Println("Senders Account not found")
		return false
	}

	acc := Accounts[accNo]
	if tacc, exists := Accounts[(transferAccNo)]; exists {
		if acc.availBal >= amount {
			// Deduct from sender
			acc.availBal -= amount
			Accounts[accNo] = acc
			// Add to receiver
			tacc.availBal += amount
			Accounts[transferAccNo] = tacc
			return true
		} else {
			fmt.Println("Incorrect pin or insufficient account balance")
			return false
		}
	} else {
		fmt.Println("Receivers Account not found")
		return false
	}
}

func Deposit(accNo, pin, amount int ) bool {
	if !Authenticate(accNo, pin) {
		return false
	}

	acc := Accounts[accNo]
	acc.availBal += amount
	Accounts[accNo] = acc
	return true
}

func Withdraw(accNo, pin, amount int) bool {
	if !Authenticate(accNo, pin) {
		return false
	}

	acc := Accounts[accNo]
	if acc.availBal >= amount {
		acc.availBal -= amount
		Accounts[accNo] = acc
		return true
	} else {
		fmt.Println("Insufficient funds")
		return false
	}
}

func Authenticate(accNo , pin int) bool {
	acc, exists := Accounts[accNo]
	if !exists {
		fmt.Println("Incorrect Account Number")
		return false
	}
	if acc.pinNo != pin {
		fmt.Println("Incorrect PIN")
		return false
	}

	return true
}

func CheckAccountBalance(accNo , pin int) int {
	if Authenticate(accNo, pin) {
		return int(Accounts[accNo].availBal)
	} else {
		return -1
	}
}

func CreateAccounts(accs []Account) {
	for _, account := range accs {
		Accounts[account.accNo] = account
	}
}

var (
	a1 = Account{
		accNo:    (rand.Int()),
		availBal: 5000,
		accType:  savingsAccount,
		pinNo:    1234,
	}

	a2 = Account{
		accNo:    (rand.Int()),
		availBal: 50001,
		accType:  currentAccount,
		pinNo:    8493,
	}

	a3 = Account{
		accNo:    (rand.Int()),
		availBal: 15000,
		accType:  savingsAccount,
		pinNo:    2474,
	}

	a4 = Account{
		accNo:    (rand.Int()),
		availBal: 1000,
		accType:  currentAccount,
		pinNo:    1927,
	}

	a5 = Account{
		accNo:    (rand.Int()),
		availBal: 500000,
		accType:  savingsAccount,
		pinNo:    5050,
	}
)

func init() {
	var accs []Account
	accs = append(accs, a1, a2, a3, a4, a5)
	CreateAccounts(accs)
}

func main() {
	fmt.Println(CheckAccountBalance(a1.accNo, 1234))
	Transfer(a1.accNo, 1234, a2.accNo, 1000)
	fmt.Println(CheckAccountBalance(a1.accNo, 1234))
	Withdraw(a1.accNo, 1234, 1500)
	fmt.Println(CheckAccountBalance(a1.accNo, 1234))
	Deposit(a1.accNo, 1234, 15000)
	fmt.Println(CheckAccountBalance(a1.accNo, 1234))

	// Wrong pin test
	fmt.Println(CheckAccountBalance(a1.accNo, 1236))

	// Wrong account number test
	fmt.Println(CheckAccountBalance(a1.accNo+5, 1234))
	
}
