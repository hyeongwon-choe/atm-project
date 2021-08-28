package atm

import (
	"testing"

	"github.com/hyeongwon-choe/atm-project/dao/dummy_dao"
	"github.com/stretchr/testify/require"
)

func Test_Account_Basic_Scenario(t *testing.T) {
	var dao dummy_dao.DummyDao
	atm := NewAtm(&dao)

	card, err := atm.PutCard(dummy_dao.CARD_NUMBER)
	require.Nil(t, err)

	accounts, err := card.GetAccounts(dummy_dao.PIN_NUMBER)
	require.Nil(t, err)

	var expectedAccounts []*Account
	for _, accountNumber := range dummy_dao.Accounts {
		expectedAccounts = append(expectedAccounts, newAccount(&dao, accountNumber))
	}
	require.Equal(t, expectedAccounts, accounts)

	balance, err := accounts[0].GetBalance()
	require.Nil(t, err)

	depositeMoney := 10
	require.Nil(t, accounts[0].Deposit(depositeMoney))

	afterBalance, err := accounts[0].GetBalance()
	require.Nil(t, err)
	require.Equal(t, balance+depositeMoney, afterBalance)

	withdrawMoney := 5
	require.Nil(t, accounts[0].Withdraw(withdrawMoney))

	finalBalance, err := accounts[0].GetBalance()
	require.Nil(t, err)
	require.Equal(t, afterBalance-withdrawMoney, finalBalance)
}
