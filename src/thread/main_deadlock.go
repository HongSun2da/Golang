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
// 	a.balance -= val
// }

// // 임금
// func (a *Account) Deposit(val int) {
// 	a.balance += val
// }

// // 잔금
// func (a *Account) Balance() int {
// 	balance := a.balance
// 	return balance
// }

// var accounts []*Account

// func Transfer(sender, receiver int, money int) {
// 	accounts[sender].mutex.Lock()
// 	accounts[receiver].mutex.Lock()

// 	accounts[sender].Widthdraw(money)
// 	accounts[receiver].Deposit(money)

// 	accounts[receiver].mutex.Unlock()
// 	accounts[sender].mutex.Unlock()

// 	fmt.Println("Transfer", sender, receiver, money)

// }

// // 계좌별 총 잔액
// func GetTotalBalance() int {
// 	total := 0
// 	for i := 0; i < len(accounts); i++ {
// 		total += accounts[i].Balance()
// 		//fmt.Printf("계좌(%d) Total: %d\n", i, total)
// 	}
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

// func main_deadlock() {
// 	for i := 0; i < 20; i++ {
// 		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
// 		//fmt.Printf("계좌생성(%d): %d\n", i, accounts[i].Balance())
// 	}

// 	go func() {
// 		for {
// 			Transfer(0, 1, 100)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			Transfer(1, 0, 100)
// 		}
// 	}()

// 	for {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println("time.Sleep -----------------------------------------")
// 	}
// }
