// 2043. 简易银行系统
// URL：https://leetcode-cn.com/problems/simple-bank-system/
// 难度：中等
// 关键词：设计、模拟
// 执行用时：316 ms, 在所有 Go 提交中击败了 91.30% 的用户
// 内存消耗：36.9 MB, 在所有 Go 提交中击败了 34.78% 的用户

package main

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance}
}

func (b *Bank) isExist(account int) bool {
	return 1 <= account && account <= len(b.balance)
}

func (b *Bank) isEnough(account int, money int64) bool {
	return b.balance[account-1] >= money
}

func (b *Bank) increase(account int, money int64) {
	b.balance[account-1] += money
}

func (b *Bank) decrease(account int, money int64) {
	b.balance[account-1] -= money
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !(b.isExist(account1) && b.isExist(account2) && b.isEnough(account1, money)) {
		return false
	}
	b.decrease(account1, money)
	b.increase(account2, money)
	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if !b.isExist(account) {
		return false
	}
	b.increase(account, money)
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if !(b.isExist(account) && b.isEnough(account, money)) {
		return false
	}
	b.decrease(account, money)
	return true
}
