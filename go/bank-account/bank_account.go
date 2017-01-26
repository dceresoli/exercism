package account

import (
	"sync"
)

type Account struct {
	sync.Mutex
	balance int64
	active  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, active: true}
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.active {
		return a.balance, true
	}
	return 0, false
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.active {
		a.active = false
		return a.balance, true
	}
	return 0, false
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.active {
		if a.balance+amount < 0 {
			return 0, false
		}
		a.balance += amount
		return a.balance, true
	}
	return 0, false
}
