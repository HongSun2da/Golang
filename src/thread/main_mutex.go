package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type Account struct {
// 	balance int
// 	mutex   *sync.Mutex
// }

// // 출금
// func (a *Account) Widthdraw(val int) {
// 	a.mutex.Lock()
// 	a.balance -= val
// 	a.mutex.Unlock()
// }

// // 임금
// func (a *Account) Deposit(val int) {
// 	a.mutex.Lock()
// 	a.balance += val
// 	a.mutex.Unlock()
// }

// // 잔금
// func (a *Account) Balance() int {
// 	a.mutex.Lock()
// 	balance := a.balance
// 	a.mutex.Unlock()
// 	return balance

// }

// var accounts []*Account
// var globalLock *sync.Mutex

// func Transfer(sender, receiver int, money int) {
// 	globalLock.Lock()
// 	//time.Sleep(100 * time.Millisecond)
// 	//fmt.Printf("Transfer : 계좌(%d) -> 계좌(%d) Money: %d\n", sender, receiver, money)
// 	accounts[sender].Widthdraw(money)
// 	//fmt.Printf("Transfer : 계좌(%d) Balance: %d\n", sender, accounts[sender].balance)
// 	accounts[receiver].Deposit(money)
// 	//fmt.Printf("Transfer : 계좌(%d) Balance: %d\n", receiver, accounts[receiver].balance)
// 	globalLock.Unlock()
// }

// // 계좌별 총 잔액
// func GetTotalBalance() int {
// 	globalLock.Lock()
// 	total := 0
// 	for i := 0; i < len(accounts); i++ {
// 		total += accounts[i].Balance()
// 		//fmt.Printf("계좌(%d) Total: %d\n", i, total)
// 	}
// 	globalLock.Unlock()
// 	return total
// }

// func RandomTranfer() {
// 	var sender, balance int
// 	for {
// 		sender = rand.Intn(len(accounts))
// 		balance = accounts[sender].Balance()
// 		if balance > 0 {
// 			break
// 		}
// 	}

// 	var receiver int
// 	for {
// 		receiver = rand.Intn(len(accounts))
// 		if sender != receiver {
// 			break
// 		}
// 	}

// 	money := rand.Intn(balance)
// 	Transfer(sender, receiver, money)
// }

// func GoTransfer() {
// 	for {
// 		RandomTranfer()
// 	}
// }

// func PrintTotalBalance() {
// 	fmt.Printf("Total:  %d\n", GetTotalBalance())
// }

// func main_mutex() {
// 	for i := 0; i < 20; i++ {
// 		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
// 		fmt.Printf("계좌생성(%d): %d\n", i, accounts[i].Balance())
// 	}

// 	globalLock = &sync.Mutex{}

// 	PrintTotalBalance()

// 	// thread 생성
// 	for i := 0; i < 10; i++ {
// 		go GoTransfer()
// 	}

// 	for {
// 		PrintTotalBalance()
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }
