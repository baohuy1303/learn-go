package middleware

import(
	"errors"
	"net/http"

	"github.com/baohuy1303/learn-go/api"
	"github.com/baohuy1303/learn-go/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or token.")

//Middleware needs to take in an http.Handler
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == ""{
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return 
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken){
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return 
		}

		next.ServeHTTP(w, r)
	})
}