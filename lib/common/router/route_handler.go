package router

import (
	_ "github.com/Meruya-Technology/go-boilerplate/docs"

	usc "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	"github.com/gorilla/mux"
	hsw "github.com/swaggo/http-swagger"
)

type RouteHandler struct{}

func (r RouteHandler) Handle() *mux.Router {
	muxClient := mux.NewRouter()
	// Swagger
	muxClient.PathPrefix("/swagger").Handler(hsw.WrapHandler)
	http := muxClient.PathPrefix("/api").Subrouter()

	/// "/user/"
	userProfile := new(usc.RetrieveProfile).Execute
	http.HandleFunc("/profile", userProfile).Methods("GET")
	return muxClient
}
