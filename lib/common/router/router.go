package router

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	"github.com/gorilla/mux"
)

type Router struct{}

func (router Router) Handle() *mux.Router {

	// var dir string
	// flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()
	// r := mux.NewRouter()

	// // This will serve files under http://localhost:8000/static/<filename>
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	muxClient := mux.NewRouter()
	http := muxClient.PathPrefix("/test").Subrouter()

	/// "/user/"
	userProfile := new(usecases.RetrieveUser).Execute
	http.HandleFunc("/profile", userProfile).Methods("GET")

	/// "/user/{key}/"
	// http.HandleFunc("/{key}/", getUser).Methods("GET")
	return muxClient
}