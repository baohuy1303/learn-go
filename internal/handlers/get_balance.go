package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/baohuy1303/learn-go/api"
	"github.com/baohuy1303/learn-go/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetBalance(w http.ResponseWriter, r *http.Request){
	var params = api.Balance_Req{}
	var decoder *schema.Decoder = schema.NewDecoder()

	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil{
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.BalanceDetails
	tokenDetails = (*database).GetUserBalance(params.Username)
	if tokenDetails == nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var res = api.Balance_Res{
		Balance: tokenDetails.Balance,
		StatusCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}