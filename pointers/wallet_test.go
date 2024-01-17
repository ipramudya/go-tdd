package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("read wallet coin name", func(t *testing.T) {
		coin := Bitcoin(2)

		got := coin.GetString()
		want := "2 BTC"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("deposit some amounts", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(999)

		want := Bitcoin(999)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw some amounts", func(t *testing.T) {
		wallet := Wallet{
			balance: 999,
		}

		err := wallet.Withdraw(99)

		want := Bitcoin(900)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(100)

		assertError(t, err, ErrInsufficientFunds)

		/** pass starting balance as wanted amount, because we dont want to perform any withdraw,
		 *	so keep balance same as starting balance
		 */
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got err %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
