package main

import (
	"cofee_shop/models"
	"cofee_shop/routes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Base de données en mémoire
var drinks []models.Drink
var orders []models.Order
var orderCounter int = 1

// initialisation du menu des boissons
var menu = []models.Drink{
	{ID: "1", Name: "Espresso", Category: "coffee", BasePrice: 2.00},
	{ID: "2", Name: "Cappuccino", Category: "coffee", BasePrice: 3.00},
	{ID: "3", Name: "Lipton tea", Category: "tea", BasePrice: 9.76},
	{ID: "4", Name: "Iced cofee", Category: "Iced cofee", BasePrice: 10.99},
	{ID: "5", Name: "Mocha", Category: "coffee", BasePrice: 4.50},
	{ID: "6", Name: "Latte", Category: "coffee", BasePrice: 3.50},
	{ID: "7", Name: "Americano", Category: "coffee", BasePrice: 2.50},
	{ID: "8", Name: "Green tea", Category: "tea", BasePrice: 8.00},
	{ID: "9", Name: "Hot chocolate", Category: "chocolate", BasePrice: 4.00},
	{ID: "10", Name: "Flat White", Category: "coffee", BasePrice: 3.75},
}

// GEt menu -recuperer le menu des boissons
func Getmenu(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)

}

// GEt /menu/{id} - recuperer une boissons spécifique par son ID
func GetDrink(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2. recuperer l'id de la boisson depuis l'url
	vars := mux.Vars(r)
	id := vars["id"]
	//3 parcourir la liste des boissons pour trouver celle qui correspond
	for _, drink := range menu {
		//4 si la boisson est trouvée avce son ID, on retourne la boissons
		if drink.ID == id {
			json.NewEncoder(w).Encode(drink)
			return
		}
	}
	//5 si la boisson n'est pas trouvée... affiche un message
	http.Error(w, "boisson non trouvée", http.StatusNotFound)
}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := routes.Register_routes()

	println("démarrage du serveur sur le port 8080....")
	println("Bienvenue très cher")
	println("J'espère que vous avez soif!!!")
	http.ListenAndServe(":8080", corsMiddleware(r))

}
