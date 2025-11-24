package handlers

import (
	"cofee_shop/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// creation du menu des boissons
var Drinks = []models.Drink{
	{ID: "1", Name: "Espresso", Category: "coffee", BasePrice: 2.00},
	{ID: "2", Name: "Cappuccino", Category: "coffee", BasePrice: 3.00},
	{ID: "3", Name: "Lipton tea", Category: "tea", BasePrice: 9.76},
	{ID: "4", Name: "Iced cofee", Category: "Iced cofee", BasePrice: 10.99},
	{ID: "5", Name: "Latte", Category: "coffee", BasePrice: 4.00},
	{ID: "6", Name: "Mocha", Category: "coffee", BasePrice: 4.50},
	{ID: "7", Name: "Green tea", Category: "tea", BasePrice: 8.50},
	{ID: "8", Name: "Black tea", Category: "tea", BasePrice: 7.50},
}

// GEt menu -recuperer le menu des boissons
func Getmenu(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")

	//3. encoder le menu en json et l'envoyer dans la reponse
	json.NewEncoder(w).Encode(Drinks)
}

// GEt /menu/{id} - recuperer une boissons spécifique par son ID
func GetDrink(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2. recuperer l'id de la boisson depuis l'url
	vars := mux.Vars(r)
	id := vars["id"]
	//3 parcourir la liste des boissons pour trouver celle qui correspond
	for _, drink := range Drinks {
		//4 si la boisson est trouvée avce son ID, on retourne la boissons
		if drink.ID == id {
			json.NewEncoder(w).Encode(drink)
			return
		}
	}
	//5 si la boisson n'est pas trouvée... affiche un message
	http.Error(w, "boisson non trouvée", http.StatusNotFound)
}
