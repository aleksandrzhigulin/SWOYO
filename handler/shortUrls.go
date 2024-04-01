package handler

import (
	"SWOYO/services"
	"context"
	"github.com/go-chi/chi"
	"net/http"
)

func shortUrls(router chi.Router) {
	router.Post("/", services.CreateShortUrl)

	router.Route("/{url}", func(router2 chi.Router) {
		router2.Use(UrlContext)
		router2.Get("/", services.GetFullUrl)
	})
}

func UrlContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := chi.URLParam(r, "url")
		ctx := context.WithValue(r.Context(), "url", url)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
