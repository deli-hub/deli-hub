package accounts

// Account struct
type Account struct {
	// 변수명이 대문자로 시작할 경우는 export가 가능해서 public / 소문자면 private
	owner   string
	balance int
}

// NewAccount creates Account (How to make a contructor)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 주소값을 복사해서 내보낸다. (반환값이 복사본이 아닌 객체임)
	return &account
}
