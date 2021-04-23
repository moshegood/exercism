package account

import (
	"sync"
)

type Account struct {
	balance int
	open    bool
	lock    sync.Mutex
}

func Open(amt int) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{balance: amt, open: true}
}

func (a *Account) Balance() (int, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.balance, a.open
}

func (a *Account) Close() (int, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if !a.open {
		return 0, false
	}
	b := a.balance
	a.balance = 0
	a.open = false
	return b, true
}

func (a *Account) Deposit(amt int) (int, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	newTotal := a.balance + amt
	if newTotal < 0 {
		return a.balance, false
	}
	a.balance = a.balance + amt
	return a.balance, a.open
}
