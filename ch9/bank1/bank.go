// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
type withrawRequest struct {
	amount   int
	response chan bool
}

var withraw = make(chan withrawRequest) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withraw(amount int) bool {
	response := make(chan bool)
	withraw <- withrawRequest{amount: amount, response: response}
	return <-response
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case request := <-withraw:
			if balance >= request.amount {
				balance -= request.amount
				request.response <- true
			} else {
				request.response <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
