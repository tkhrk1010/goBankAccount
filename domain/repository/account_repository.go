package repository

import (
	"github.com/tkhrk1010/bank_account/domain/model"
)

// AccountRepository is an interface that defines methods for storing and retrieving accounts.
type AccountRepository interface {
	Save(a *model.Account) error
	FindById(id int) (*model.Account, error)
	Update(a *model.Account) (error)
	Delete(a *model.Account) error
}