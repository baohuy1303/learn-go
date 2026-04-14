package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	Username string
	AuthToken string
}

type BalanceDetails struct{
	Balance int64
	Username string
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetUserBalance(username string) *BalanceDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var databse DatabaseInterface = &mockDB{}

	var err error = databse.SetupDatabase()
	if err != nil{
		log.Error(err)
		return nil, err
	}

	return &databse, nil
}

