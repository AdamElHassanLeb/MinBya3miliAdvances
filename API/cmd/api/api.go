package main

import (
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Middleware"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config  config
	Service Services.Service
}

type config struct {
	address string
	db      dbConfig
}

type dbConfig struct {
	addr string
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(mainRouter chi.Router) {
		mainRouter.Route("/user", func(userRouter chi.Router) {
			userRouter.Get("/users", app.GetAllUsers)
			userRouter.Get("/userId/{id}", app.GetUserById)
			userRouter.Get("/userName/{name}", app.GetUserByName)
			userRouter.Post("/create", app.CreateUser)
			userRouter.With(Middleware.AuthMiddleware).Delete("/delete/{id}", app.DeleteUser)
			userRouter.With(Middleware.AuthMiddleware).Put("/update/{id}", app.UpdateUser)
			userRouter.Get("/auth", app.authUser)
		})
		mainRouter.Route("/listing", func(listingRouter chi.Router) {
			listingRouter.Get("/listings/{type}", app.GetAllListings)
			listingRouter.Get("/listingId/{id}", app.GetListingByID)
			listingRouter.Get("/listings/user/{user_id}/{type}", app.GetListingsByUserID)
			listingRouter.Get("/search/{query}/{type}", app.GetListingsBySearch)
			listingRouter.Get("/date/{type}", app.GetListingsByDate)
			listingRouter.Get("/date/search/{query}/{type}", app.GetListingsByDateAndSearch)
			listingRouter.Get("/distance/{latitude}/{longitude}/{max_distance}/{type}", app.GetListingsByDistance)
			listingRouter.Get("/location/{latitude}/{longitude}/{max_range}/{type}", app.GetListingsByLocation)
			listingRouter.Post("/create", app.CreateListing)
			listingRouter.Put("/update/{id}", app.UpdateListing)
			listingRouter.Delete("/delete/{id}", app.DeleteListing)
		})

	})

	return r
}

func (app *application) run(mux http.Handler) error {

	server := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("starting server at %s", app.config.address)

	return server.ListenAndServe()
}
