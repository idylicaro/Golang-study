package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		verifyBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		verifyBalance(t, wallet, Bitcoin(10))
		verifyErrorNonexistent(t, err)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		verifyBalance(t, wallet, Bitcoin(20))
		verifyHasError(t, err, ErrInsufficientFunds)
	})
}

func verifyBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}

func verifyErrorNonexistent(t *testing.T, result error) {
	t.Helper()
	if result != nil {
		t.Fatal("didn't expect an error but got one")
	}

}
func verifyHasError(t *testing.T, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("expected an error but didn't get one")
	}
	if result != expected {
		t.Errorf("got %s expected %s", result, expected)
	}

}
