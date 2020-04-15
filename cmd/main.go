package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/tomoyamachi/chi-oapi/pkg/api"
	"github.com/tomoyamachi/chi-oapi/pkg/gen/store"
	"github.com/tomoyamachi/chi-oapi/pkg/gen/user"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	storeService := api.StoreService{}
	userService := api.UserService{}
	s := store.HandlerFromMux(storeService, r)
	r = s.(*chi.Mux)
	s = user.HandlerFromMux(userService, r)

	// Mount the admin sub-router, which btw is the same as:
	// r.Route("/admin", func(r chi.Router) { admin routes here })

	// Passing -routes to the program will generate docs for the above
	// router definition. See the `routes.json` file in this folder for
	// the output.
	http.ListenAndServe(":3333", s)
}
