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

	r.Route("/api", func(v1Router chi.Router) {
		v1Router.Route("/v1", func(mainRouter chi.Router) {
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
				listingRouter.Get("/distance/{longitude}/{latitude}/{max_distance}/{type}", app.GetListingsByDistance)
				//listingRouter.Get("/location/{longitude}/{latitude}/{max_range}/{type}", app.GetListingsByLocation)
				listingRouter.With(Middleware.AuthMiddleware).Post("/create", app.CreateListing)
				listingRouter.With(Middleware.AuthMiddleware).Put("/update/{id}", app.UpdateListing)
				listingRouter.With(Middleware.AuthMiddleware).Delete("/delete/{id}", app.DeleteListing)
			})

			mainRouter.Route("/image", func(imageRouter chi.Router) {
				imageRouter.With(Middleware.AuthMiddleware).Post("/upload/{listing_id}", app.createImage)          // Upload image
				imageRouter.With(Middleware.AuthMiddleware).Get("/imageId/{image_id}", app.GetImageByID)           // Get image by ID
				imageRouter.With(Middleware.AuthMiddleware).Get("/listing/{listing_id}", app.GetImagesByListingID) // Get images by listing ID
				imageRouter.With(Middleware.AuthMiddleware).Get("/user/{user_id}", app.GetImagesByUserID)          // Get images by user ID
				imageRouter.With(Middleware.AuthMiddleware).Get("/profile/{user_id}", app.GetImagesByUserProfile)  // Get images by user with profile set to true
				imageRouter.With(Middleware.AuthMiddleware).Delete("/{image_id}", app.DeleteImage)                 // Delete image
				imageRouter.With(Middleware.AuthMiddleware).Put("/{image_id}/{show_on_profile}", app.UpdateImage)  // Update image
			})

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
