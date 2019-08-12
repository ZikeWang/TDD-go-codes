package main

import "testing"

func TestWallet(t *testing.T) {
	//wallet := Wallet{}
	//each t.Run tests should be seperated, so it's better to initialize seperately

	checkequal := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("test deposit and balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		checkequal(t, got, want)
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(7))
		got := wallet.Balance()
		want := Bitcoin(13)
		checkequal(t, got, want)
	})
}
