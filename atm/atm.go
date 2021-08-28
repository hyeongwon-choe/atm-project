package atm

import "github.com/hyeongwon-choe/atm-project/dao"

type Atm struct {
	dao dao.Dao
}

func NewAtm(dao dao.Dao) *Atm {
	return &Atm{dao}
}

func (atm *Atm) PutCard(cardNumber string) (*Card, error) {
	err := atm.dao.CheckCardNumber(cardNumber)
	if err != nil {
		if err == dao.ErrNotFound {
			return nil, ErrInvalidCardNumber
		}
		return nil, err
	}

	return newCard(atm.dao, cardNumber), nil
}
