package leetcode

type Bank []int64

func Constructor(balance []int64) Bank {
	b := make([]int64, len(balance))
	copy(b, balance)
	return b
}

func (b Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !b.validAccount(account1) || !b.validAccount(account2) || b[account1-1] < money {
		return false
	}
	b[account1-1] -= money
	b[account2-1] += money
	return true
}

func (b Bank) Deposit(account int, money int64) bool {
	if !b.validAccount(account) {
		return false
	}
	b[account-1] += money
	return true
}

func (b Bank) Withdraw(account int, money int64) bool {
	if !b.validAccount(account) || b[account-1] < money {
		return false
	}
	b[account-1] -= money
	return true
}

func (b Bank) validAccount(account int) bool {
	return account > 0 && account <= len(b)
}
