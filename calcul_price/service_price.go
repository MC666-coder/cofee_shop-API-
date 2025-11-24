package calcul_price

// fonction pour calculer le prix total de la commande
func CalculatePrice(basePrice float64, size string, extras []string) float64 {
	//ajouter selon la taille
	//small: *0.8
	switch size {
	case "small":
		basePrice *= 0.8
	case "medium":
		basePrice *= 1.0
	case "large":
		basePrice *= 1.3
	}
	//3 ajouter 0.5 par extra
	for range len(extras) {
		basePrice += 0.5
	}
	return basePrice
}
