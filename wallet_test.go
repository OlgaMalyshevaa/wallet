package wallet

import(
	"testing"
	"fmt"
)

func TestWallet (t *testing.T) {

	assertWallet := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
	
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("not enough")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoErrors := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("qw")
		}
	}

	t.Run("deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))	
		fmt.Printf("address wallet is: %p\n", &wallet.balance)
		
		assertWallet(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoErrors(t, err)
		assertWallet(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(40))

		assertWallet(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})

}