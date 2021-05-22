package main

import (
	"fmt"

	"github.com/deli-hub/learningGo/accounts"
)

func main() {
	account := accounts.NewAccount("merry")
	// 아래 코드는 작동하지 않는다. balance가 private기 때문!
	// account.balance = 10000
	account.Deposit(1000)
	fmt.Println(account)
	// err := account.Withdraw(1500)
	// if err != nil {
	// 	// 로그를 출력하고 프로그램을 종료한다.
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
}
