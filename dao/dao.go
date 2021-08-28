package dao

import "fmt"

var (
	ErrNotFound        = fmt.Errorf("ErrNotFound")
	ErrInvalidArgument = fmt.Errorf("ErrInvalidArgument")
	ErrInvalidPin      = fmt.Errorf("ErrInvalidPin")
)

type Dao interface {
	CheckCardNumber(cardNumber string) error
	CheckPinNumber(cardNumber, pin string) error
	GetAccounts(cardNumber string) ([]string, error)
	GetBalance(account string) (int, error)
	SetBalance(account string, balance int) error
}
