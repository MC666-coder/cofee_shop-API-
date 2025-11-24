package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cofee_shop/models"

	"github.com/gorilla/mux"
)

// Base de données en mémoire
var drinks []models.Drink
var orders []models.Order
var orderCounter int = 1

// GEt menu -recuperer le menu des boissons
func getmenu(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")

	//2.creation du menu des boissons
	menu := []models.Drink{
		{ID: "1", Name: "Espresso", Category: "coffee", BasePrice: 2.00},
		{ID: "2", Name: "Cappuccino", Category: "coffee", BasePrice: 3.00},
		{ID: "3", Name: "Lipton tea", Category: "tea", BasePrice: 9.76},
		{ID: "4", Name: "Iced cofee", Category: "Iced cofee", BasePrice: 10.99},
	}
	//3. encoder le menu en json et l'envoyer dans la reponse
	json.NewEncoder(w).Encode(menu)
}

// GEt /menu/{id} - recuperer une boissons spécifique par son ID
func getDrink(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2. recuperer l'id de la boisson depuis l'url
	vars := mux.Vars(r)
	id := vars["id"]
	menu := []models.Drink{
		{ID: "1", Name: "Espresso", Category: "coffee", BasePrice: 2.00},
		{ID: "2", Name: "Cappuccino", Category: "coffee", BasePrice: 3.00},
		{ID: "3", Name: "Lipton tea", Category: "tea", BasePrice: 9.76},
		{ID: "4", Name: "Iced cofee", Category: "Iced cofee", BasePrice: 10.99},
	}
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

// post /orders - créer une nouvelle commande
func createOrder(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2 decoder le body JSOn dans une variable de type order
	var neworder models.Order
	//3 gérer les erreurs de décodage
	err := json.NewDecoder(r.Body).Decode(&neworder)
	if err != nil {
		http.Error(w, "décodage de la commande a échoué", http.StatusBadRequest)
	}
	//4 verifier l'existance de la boissons commandée
	var comm_boissons *models.Drink
	for _, drinks := range drinks {
		if drinks.ID == neworder.DrinkID {
			comm_boissons = &drinks
			break
		}
	}
	//5 si la boisson n'existe pas, retourner une erreur
	if comm_boissons == nil {
		http.Error(w, "404 Bad request", http.StatusBadRequest)
	}
	//6 génerer un Id pour la commande
	neworder.ID = fmt.Sprintf("ORD-%03d", orderCounter)
	orderCounter++
	//7 remplir le order avec le nom de la boisson
	neworder.DrinkName = comm_boissons.Name
	//8 definir order statu à sattus pending
	neworder.Status = models.StatusPending
	//9 definir order.orderedAt à time.now()
	neworder.OrderedAt = time.Now()
	//10 calculer le prix total (appeler calculateprice)
	neworder.TotalPrice = calculatePrice(comm_boissons.BasePrice, neworder.Size, neworder.Extras)
	//11 ajouter la nouvelle commande à la liste des commandes
	orders = append(orders, neworder)
	//12 retounrer 201 created avec la commande JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(neworder)
}

// fonction pour calculer le prix total de la commande
func calculatePrice(basePrice float64, size string, extras []string) float64 {

	//partir du baseprix de la boisson

	return basePrice
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
	r := mux.NewRouter()
	r.HandleFunc("/menu", getmenu).Methods("GET")
	r.HandleFunc("/menu{id}", getDrink).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	println("démarrage du serveur sur le port 8080....")
	http.ListenAndServe(":8080", corsMiddleware(r))

}
