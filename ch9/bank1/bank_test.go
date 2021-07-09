// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/bank1"
)

func TestBank(t *testing.T) {
	for i := 0; i < 10; i++ {
		done := make(chan bool)

		// Alice
		go func() {
			bank.Deposit(200)
			fmt.Println("=", bank.Balance())
			done <- false
		}()

		// Bob
		go func() {
			bank.Deposit(100)
			done <- false
		}()

		// Zé
		go func() {
			withrawed := bank.Withraw(100)
			fmt.Println("withrawed: ", withrawed)
			done <- withrawed
		}()

		// Wait for both transactions.
		result := <-done
		result = result || <-done
		result = result || <-done
		expected := 0
		if result {
			expected = 200
		} else {
			expected = 300
		}

		if got, want := bank.Balance(), expected; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
		fmt.Println(bank.Withraw(expected))
	}
}
