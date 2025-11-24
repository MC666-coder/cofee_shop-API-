package handlers

import (
	calculprice "cofee_shop/calcul_price"
	"cofee_shop/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Base de données en mémoire
var orders []models.Order
var orderCounter int = 1

// post /orders - créer une nouvelle commande
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2 decoder le body JSOn dans une variable de type order
	var neworder models.Order
	//3 gérer les erreurs de décodage
	err := json.NewDecoder(r.Body).Decode(&neworder)
	if err != nil {
		http.Error(w, "décodage de la commande a échoué", http.StatusBadRequest)
		return
	}
	//4 verifier l'existance de la boissons commandée
	var comm_boissons *models.Drink
	for i := range Drinks {
		if Drinks[i].ID == neworder.DrinkID {
			comm_boissons = &Drinks[i]
			break
		}
	}
	/*for _, drinks := range Drinks {
		if drinks.ID == neworder.DrinkID {
			comm_boissons = &drinks
			break
		}
	}*/
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
	//10 calculer le prix total (appeler CalculatePrice du package calcul_price)
	neworder.TotalPrice = calculprice.CalculatePrice(comm_boissons.BasePrice, neworder.Size, neworder.Extras)
	//11 ajouter la nouvelle commande à la liste des commandes
	orders = append(orders, neworder)
	//12 retounrer 201 created avec la commande JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(neworder)
}

// Get /orders - récuperer toutes les commandes
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2. encoder et retourner la liste orders en Json
	json.NewEncoder(w).Encode(orders)
}

// Get /Oders/{id} - récuperer une commande spécifique
func Getorder(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2. recuperer l'id de la commande depuis les varibles de route
	vars := mux.Vars(r)
	id := vars["id"]
	//3. parcourir la liste orders
	for _, order := range orders {
		//4 si trouvé : encoder et retourner le commande
		if order.ID == id {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	//5 sinon : retouner une erreur 404
	http.Error(w, "erreur 404", http.StatusNotFound)
}

// PATH /orders/{id} - changer le statut d'une commande
func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	//1. definir le content type à l'apllication json
	w.Header().Set("Content-Type", "application/json")
	//2. recuperer l'id depuis la varible de route
	vars := mux.Vars(r)
	id := vars["id"]
	//3. créer une structure temporaire avec un champ status
	var statusUpdate struct {
		Status models.OrderStatus `json:"status"`
	}
	//4. decoder le body JSON dans cette structure
	err := json.NewDecoder(r.Body).Decode(&statusUpdate)
	//5. gérer les erreurs de décodage
	if err != nil {
		http.Error(w, "décodage du status a échoué", http.StatusBadRequest)
		return
	}
	//6. parcourir Orders et troiuver la commande
	for i, order := range orders {
		if order.ID == id {
			//7. mettre à jour le statut de la commande
			orders[i].Status = statusUpdate.Status
			//8. retourner la commande mis à jour en JSON
			json.NewEncoder(w).Encode(orders[i])
			return
		}
	}
	//9. si non trouvée, retourner une erreur 404
	http.Error(w, "erreur 404", http.StatusNotFound)
}

// DELETE /orders/{id} - annuler une commande
func CancelOrder(w http.ResponseWriter, r *http.Request) {
	//1. recuperer l'ID depuis les varibles de route
	vars := mux.Vars(r)
	id := vars["id"]
	//2. parcourir orders avec l'index
	for i, order := range orders {
		//3. si la commande est trouvée
		if order.ID == id {
			//a. verifier que le statut n'est pas "picked-up"
			if order.Status == models.StatusPickedUp {
				http.Error(w, "404 Bad request", http.StatusBadRequest)
				return
			}
			//b. sinon suprimer le commande de la liste orders
			orders = append(orders[:i], orders[i+1:]...)
			//c. retouner un status 204 no content
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	//4. si non trouvée, retourner une erreur 404
	http.Error(w, "erreur 404", http.StatusNotFound)

}
