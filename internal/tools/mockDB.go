package tools

import(
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Huy":{
		Username: "Huy",
		AuthToken: "123",
	},
	"John":{
		Username: "John",
		AuthToken: "456",
	},
}

var mockBalanceDetails = map[string]BalanceDetails{
	"Huy":{
		Username: "Huy",
		Balance: 100,
	},
	"John":{
		Username: "John",
		Balance: 200,
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData, exist := mockLoginDetails[username]
	if !exist{
		return nil
	}
	return &clientData
}

func (db *mockDB) GetUserBalance(username string) *BalanceDetails{
	time.Sleep(1 * time.Second)
	var clientData = BalanceDetails{}
	clientData, exist := mockBalanceDetails[username]
	if !exist{
		return nil
	}
	return &clientData
}

func (db *mockDB) SetupDatabase() error{
	return nil
}