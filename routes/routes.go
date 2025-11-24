package routes

import (
	"cofee_shop/handlers"

	"github.com/gorilla/mux"
)

func Register_routes() *mux.Router {
	// créer le routeur mux
	r := mux.NewRouter()
	// définir les routes
	//GEt /menu
	//GET /menu/{id}
	r.HandleFunc("/menu", handlers.Getmenu).Methods("GET")
	//GET /menu/{id}
	r.HandleFunc("/menu/{id}", handlers.GetDrink).Methods("GET")
	//POST /orders
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	//PATCH /orders/{id}
	r.HandleFunc("/orders/{id}", handlers.UpdateOrderStatus).Methods("PATCH")
	//DELETE /orders/{id}
	r.HandleFunc("/orders/{id}", handlers.CancelOrder).Methods("DELETE")
	//GET /orders
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	//GET /orders/{id}
	r.HandleFunc("/orders/{id}", handlers.Getorder).Methods("GET")
	//GET /oders/{id}/status
	r.HandleFunc("/orders/{id}/status", handlers.Getorder).Methods("GET")

	return r

}
