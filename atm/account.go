package atm

import "github.com/hyeongwon-choe/atm-project/dao"

type Account struct {
	accountNumber string
	dao           dao.Dao
}

func newAccount(dao dao.Dao, account string) *Account {
	return &Account{account, dao}
}

func (account *Account) GetBalance() (int, error) {
	balance, err := account.dao.GetBalance(account.accountNumber)
	if err != nil {
		if err == dao.ErrNotFound {
			return 0, ErrAccountNotFound
		}
		return 0, err
	}
	return balance, nil
}

func (account *Account) Withdraw(money int) error {
	if money <= 0 {
		return ErrInvalidMoneyInput
	}

	balance, err := account.GetBalance()
	if err != nil {
		return err
	}

	balance -= money
	if balance < 0 {
		return ErrNotEnoughMoney
	}

	if err := account.dao.SetBalance(account.accountNumber, balance); err != nil {
		if err == dao.ErrNotFound {
			return ErrAccountNotFound
		}
		return err
	}
	return nil
}

func (account *Account) Deposit(money int) error {
	if money <= 0 {
		return ErrInvalidMoneyInput
	}

	balance, err := account.GetBalance()
	if err != nil {
		return err
	}

	balance += money
	if err := account.dao.SetBalance(account.accountNumber, balance); err != nil {
		if err == dao.ErrNotFound {
			return ErrAccountNotFound
		}
		return err
	}
	return nil
}

func (account *Account) GetAccountNumber() string {
	return account.accountNumber
}
