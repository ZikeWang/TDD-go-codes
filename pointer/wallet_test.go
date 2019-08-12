package main

import "testing"

func TestWallet(t *testing.T) {
	//wallet := Wallet{}
	//each t.Run tests should be seperated, so it's better to initialize seperately

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("test deposit and balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(7))
		want := Bitcoin(13)
		assertBalance(t, wallet, want)
	})
}
