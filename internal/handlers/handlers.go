package handlers

import(
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/baohuy1303/learn-go/internal/middleware"
)

func Handler(r *chi.Mux){
	//GLOBAL MIDDLEWARE
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router){
		//Middleware first
		router.Use(middleware.Authorization)

		// Then get
		router.Get("/balance", GetBalance)
	})
}