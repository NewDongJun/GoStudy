package main

import (
	"fmt"

	"github.com/newdongjun/GoStudy/banking"
)

func main() {
	account := banking.NewAccount("dong")
	fmt.Println(account)
}
