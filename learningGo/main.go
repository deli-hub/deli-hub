package main

import (
	"fmt"

	"github.com/deli-hub/learningGo/accounts"
)

func main() {
	account := accounts.NewAccount("merry")
	// 아래 코드는 작동하지 않는다. balance가 private기 때문!
	// account.balance = 10000
	fmt.Println(account)
}
