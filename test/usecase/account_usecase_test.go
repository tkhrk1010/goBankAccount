package usecase

import (
	// "errors"
	"testing"

	"github.com/tkhrk1010/go_bank_account/src/usecase"
	"github.com/tkhrk1010/go_bank_account/src/infrastructure/repository"
)

func TestAccountUsecase_OpenAccount(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()
	usecase := usecase.NewAccountUsecase(repo)

	account, err := usecase.OpenAccount(100)
	if err != nil {
			t.Errorf("error opening account: %s", err.Error())
	}

	if account == nil {
			t.Error("account should not be nil")
	}

	if account.Id <= 0 {
			t.Error("account ID should be greater than zero")
	}

	if account.Balance != 100 {
			t.Errorf("account balance should be 100 but got %d", account.Balance)
	}

}
