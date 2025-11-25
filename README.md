☕ Coffee Shop API avec Go & Gorilla Mux

Une API simple permettant de gérer un coffee shop : consulter le menu, passer des commandes, suivre leur statut et les annuler.
Ce projet est réalisé en Go avec le framework Gorilla Mux et utilise un stockage en mémoire.

---

🚀 Fonctionnalités
- 📋 Menu : consulter toutes les boissons disponibles ou une boisson spécifique
- 📝 Commandes : créer une nouvelle commande
- 📦 Suivi : lister toutes les commandes ou en consulter une
- 🔄 Mise à jour : changer le statut d’une commande (pending → ready → picked-up)
- ❌ Annulation : supprimer une commande si elle n’a pas encore été récupérée

---

🛠️ Technologies
- Langage : Go
- Framework : Gorilla Mux

---

📂 Installation & Setup

1. Cloner le projet
```bash
git clone https://github.com/MC666-coder/cofee_shop-API-.git
cd coffee-shop-api
```

2. Initialiser le module Go
```bash
go mod init coffee-shop-api
go get -u github.com/gorilla/mux
```

3. Lancer le serveur
```bash
air
```

Le serveur démarre sur http://localhost:8080

---

📑 Endpoints disponibles

Menu
- GET /menu → retourne toutes les boissons
- GET /menu/{id} → retourne une boisson spécifique

Commandes
- POST /orders → créer une commande
- GET /orders → lister toutes les commandes
- GET /orders/{id} → consulter une commande spécifique
- PATCH /orders/{id}/status → mettre à jour le statut d’une commande
- DELETE /orders/{id} → annuler une commande

---

🧪 Exemples de requêtes (via curl)

Voir le menu
```bash
curl http://localhost:8080/menu
```

Passer une commande
```bash
curl -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-d '{
  "drink_id": "2",
  "size": "large",
  "extras": ["milk", "sugar"],
  "customer_name": "Alice"
}'
```

Changer le statut
```bash
curl -X PATCH http://localhost:8080/orders/ORD-001/status \
-H "Content-Type: application/json" \
-d '{"status": "ready"}'
```

Annuler une commande
```bash
curl -X DELETE http://localhost:8080/orders/ORD-001
```

---

🌐 Interface Web de test
Une interface simple est disponible pour tester l’API :
```bash
 https://hellodamien.github.io/drink-ordering-app/
```

---

📌 Notes
- Les données sont stockées en mémoire (elles disparaissent à l’arrêt du serveur).
- Les IDs de commandes sont générés automatiquement (ORD-001, ORD-002, etc.).
- Le middleware CORS est activé pour permettre les requêtes depuis une interface web externe.

---

👨‍💻 Auteur
Projet réalisé par Claude Marvine MBOUROU(claude7.9) dans le cadre d’un TP Go.
Vous pouvez maintenant admirer mon talent (en toute modestie bien sûr :) ).
