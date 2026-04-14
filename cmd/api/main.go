package main

import(
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/baohuy1303/learn-go/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil{
		log.Error(err)
	}
}