package accounts

import "errors"

// Account struct
type Account struct {
	// 변수명이 대문자로 시작할 경우는 export가 가능해서 public / 소문자면 private
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw you are poor")

// NewAccount creates Account (How to make a contructor)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 주소값을 복사해서 내보낸다. (반환값이 복사본이 아닌 객체임)
	return &account
}

// Deposit x amount on your account
/* func (a Account) Deposit(amount int) {
   여기서 받아오는 receiver Account는 struct의 복사본이다.
   즉, 복사본의 balance만 바꿔줬기 때문에 원본에서는 적용이 되지 않는다. */
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your amount
// errors handling
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
