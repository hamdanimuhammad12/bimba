package routes

import (
	"bimba/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/belanja", homeGetHandler).Methods("GET")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/", middleware.AuthRequired(adminGetHandler)).Methods("GET")
	r.HandleFunc("/logout", middleware.AuthRequired(logoutGetHandler)).Methods("GET")
	r.HandleFunc("/dashboard", middleware.AuthRequired(dashboardGetHandler)).Methods("GET")
	r.HandleFunc("/aaaa", middleware.AuthRequired(kwitanGetHandler)).Methods("GET")

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
