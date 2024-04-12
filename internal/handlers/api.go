package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/middleware"
	"github.com/heyitschauncey/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
