package routes

import (
	"wayshub-server/handlers"
	"wayshub-server/pkg/middleware"
	"wayshub-server/pkg/mysql"
	"wayshub-server/repositories"

	"github.com/gorilla/mux"
)

func SubscribeRoutes(r *mux.Router) {
	subscribeRepository := repositories.RepositorySubscribe(mysql.DB)
	h := handlers.HandlerSubscribe(subscribeRepository)

	r.HandleFunc("/subscribe", middleware.Auth(h.Subscribe)).Methods("POST")
	r.HandleFunc("/subscribe/{id}", middleware.Auth(h.Subscribe)).Methods("DELETE")
	r.HandleFunc("/subscribe", middleware.Auth(h.Subscription)).Methods("GET")
	
}
