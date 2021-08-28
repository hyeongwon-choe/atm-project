package atm

import "fmt"

var (
	ErrNotFound        = fmt.Errorf("ErrNotFound")
	ErrInvalidArgument = fmt.Errorf("ErrInvalidArgument")
)

type AtmDao interface {
	GetPinNumber(cardNumber string) (string, error)
}

type CardDao interface {
	GetAccounts(cardNumber string) ([]string, error)
}

type AccountDao interface {
	GetBalance(account string) (int, error)
	SetBalance(account string, balance int)
}
