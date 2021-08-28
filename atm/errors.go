package atm

import "fmt"

var (
	ErrInvalidCardNumber = fmt.Errorf("ErrInvalidCardNumber")
	ErrInvalidPin        = fmt.Errorf("ErrInvalidPin")
	ErrNotEnoughMoney    = fmt.Errorf("ErrNotEnoughMoney")
	ErrAccountNotFound   = fmt.Errorf("ErrAccountNotFound")
	ErrInvalidMoneyInput = fmt.Errorf("ErrInvalidMoneyInput")
)
