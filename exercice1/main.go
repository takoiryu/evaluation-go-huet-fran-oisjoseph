package main

import "fmt"

func afficherMenu() {
	fmt.Println("1 - Eau       : 1 €")
	fmt.Println("2 - Soda      : 2 €")
	fmt.Println("3 - Café      : 2 €")
	fmt.Println("4 - Chocolat  : 3 €")
	fmt.Println("0 - Quitter")
}
func obtenirPrix(choix int, prix float32) float32 {
	if choix == 1 {
		prix = prix + 1
		fmt.Println("Prix :", prix, "€")
	} else if choix == 2 {
		prix = prix + 2
		fmt.Println("Prix :", prix, "€")
	} else if choix == 3 {
		prix = prix + 2
		fmt.Println("Prix :", prix, "€")
	} else if choix == 4 {
		prix = prix + 3
		fmt.Println("Prix :", prix, "€")
	}
	return prix
}

func afficherBoisson(choix int) {
	if choix == 1 {
		fmt.Println("Vous avez choisi : Eau")
	} else if choix == 2 {
		fmt.Println("Vous avez choisi : Soda")
	} else if choix == 3 {
		fmt.Println("Vous avez choisi : Café")
	} else if choix == 4 {
		fmt.Println("Vous avez choisi : Chocolat")
	}
}
func paiment(paye float32, rendu float32, demand float32, prix float32) float32 {
	if paye == prix {
		fmt.Println("Merci !")
		return 0
	} else if paye > prix {
		rendu = paye - prix
		fmt.Println("Merci !")
		fmt.Println("Votre monnaie : ", rendu, "€")
		return 0
	} else if paye < prix {
		demand = prix - paye
		fmt.Println("Montant insuffisant !")
		fmt.Println("Il manque ", demand, "€")
		return demand
	}
	return 0
}
func main() {
	var prix float32 = 0
	choix := 5
	var paye float32 = 0
	var rendu float32 = 0
	var demand float32 = 0
	afficherMenu()
	for choix != 0 {
		fmt.Printf("Votre choix :")
		fmt.Scan(&choix)
		if choix > 0 && choix <= 4 {
			afficherBoisson(choix)
			prix = obtenirPrix(choix, prix)
		} else if choix != 0 {
			fmt.Println("Choix invalide !")
		}
	}
	fmt.Printf("Montant inséré :")
	fmt.Scan(&paye)
	demand = paiment(paye, rendu, demand, prix)
	for demand != 0 {
		fmt.Printf("Montant inséré :")
		fmt.Scan(&paye)
		demand = paiment(paye, rendu, demand, demand)
	}
}
