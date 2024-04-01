package handler

import (
	"SWOYO/db"
	"github.com/go-chi/chi"
	"net/http"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.Route("/", shortUrls)
	return router
}

func NewHandlerWithoutDb() http.Handler {
	router := chi.NewRouter()
	router.Route("/", shortUrls)
	return router
}
