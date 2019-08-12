package main

import "testing"

func TestWallet(t *testing.T) {
	//wallet := Wallet{}
	//each t.Run tests should be seperated, so it's better to initialize seperately

	t.Run("test deposit and balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("test withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(7))
		want := Bitcoin(13)
		assertBalance(t, wallet, want)
		assertNoError(t, err, nil)
	})

	t.Run("test withdraw in insufficient situation", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(40))
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("want an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error, want error) {
	if got != want {
		t.Fatal("got an error but didn't want one")
	}
}
