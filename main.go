package main

import (
	"fmt"

	"github.com/newdongjun/GoStudy/banking"
)

func main() {
	account := banking.NewAccount("dong")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Banlance())
}
