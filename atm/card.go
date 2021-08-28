package atm

import "github.com/hyeongwon-choe/atm-project/dao"

type Card struct {
	cardNumber string
	dao        dao.Dao
}

func newCard(dao dao.Dao, cardNumber string) *Card {
	return &Card{cardNumber, dao}
}

func (card *Card) GetAccounts(pin string) ([]*Account, error) {
	err := card.dao.CheckPinNumber(card.cardNumber, pin)
	if err != nil {
		if err == dao.ErrInvalidPin {
			return nil, ErrInvalidPin
		}
		return nil, err
	}

	accounts, err := card.dao.GetAccounts(card.cardNumber)
	if err != nil {
		if err == dao.ErrNotFound {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	var accountArray []*Account
	for _, accountNumber := range accounts {
		accountArray = append(accountArray, newAccount(card.dao, accountNumber))
	}

	return accountArray, nil
}
