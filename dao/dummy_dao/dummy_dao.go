package dummy_dao

import "github.com/hyeongwon-choe/atm-project/dao"

type DummyDao struct{}

const (
	CARD_NUMBER = "1234-5678-1234-5678"
	PIN_NUMBER  = "1234"
)

var (
	Accounts = []string{
		"00000-00000",
		"11111-11111",
		"22222-22222",
	}
	accountsMap = map[string]int{
		"00000-00000": 0,
		"11111-11111": 10,
		"22222-22222": 20,
	}
)

func (*DummyDao) CheckCardNumber(cardNumber string) error {
	if cardNumber != CARD_NUMBER {
		return dao.ErrNotFound
	}
	return nil
}

func (*DummyDao) CheckPinNumber(cardNumber, pin string) error {
	if cardNumber != CARD_NUMBER {
		return dao.ErrNotFound
	}
	if pin != PIN_NUMBER {
		return dao.ErrInvalidPin
	}
	return nil
}

func (*DummyDao) GetAccounts(cardNumber string) ([]string, error) {
	if cardNumber != "1234-5678-1234-5678" {
		return nil, dao.ErrNotFound
	}
	return Accounts, nil
}

func (*DummyDao) GetBalance(account string) (int, error) {
	bal, ok := accountsMap[account]
	if !ok {
		return 0, dao.ErrNotFound
	}
	return bal, nil
}

func (*DummyDao) SetBalance(account string, balance int) error {
	_, ok := accountsMap[account]
	if !ok {
		return dao.ErrNotFound
	}

	if balance < 0 {
		return dao.ErrInvalidArgument
	}

	accountsMap[account] = balance

	return nil
}
