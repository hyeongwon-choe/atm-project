package dummy_dao

type DummyAtmDao struct{}

var (
	accounts = map[string]int{
		"00000-00000": 0,
		"11111-11111": 10,
		"22222-22222": 20,
	}
)

func (*DummyAtmDao) GetPinNumber(cardNumber string) (string, error) {
	if cardNumber != "1234-5678-1234-5678" {
		return "", ErrNotFound
	}
	return "1234", nil
}

func (*DummyCardDao) GetAccounts(cardNumber string) ([]string, error) {
	if cardNumber != "1234-5678-1234-5678" {
		return "", ErrNotFound
	}
	return accounts
}

func (*DummyAccoundDao) GetBalance(account string) (int, error) {
	bal, ok := accounts[account]
	if !ok {
		return 0, ErrNotFound
	}
	return bal
}

func (*DummyAccoundDao) SetBalance(account string, balance int) error {
	bal, ok := accounts[account]
	if !ok {
		return ErrNotFound
	}

	if balance < 0 {
		return ErrInvalidArgument
	}

	accounts[account] = balance

	return nil
}
