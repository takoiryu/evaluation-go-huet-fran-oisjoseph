package main

import "fmt"

func calculerMoyenne(somme int, nombre int) float64 {
	var moyenne float64
	moyenne = float64(somme) / float64(nombre)
	return moyenne
}
func trouverMaximum(lst []int) int {
	if len(lst) == 0 {
		return 0
	}
	maxi := lst[0]
	for _, n := range lst {
		if n > maxi {
			maxi = n
		}
	}
	return maxi
}
func trouverMinimum(lst []int) int {
	if len(lst) == 0 {
		return 0
	}
	mini := lst[0]
	for _, n := range lst {
		if n < mini {
			mini = n
		}
	}
	return mini
}
func afficherResultat(moyenne float64, maxi int, mini int) {
	fmt.Println("=== RÉSULTATS ===")
	fmt.Println("Moyenne :", moyenne)
	fmt.Println("Note maximale :", maxi)
	fmt.Println("Note minimale :", mini)
	if moyenne >= 10 {
		fmt.Println("L'étudiant est admis.")
	} else {
		fmt.Println("L'étudiant est non admis.")
	}
}
func main() {
	cpt := 1
	nombre := 0
	somme := 0
	note := 0
	moyenne := 0.0
	lst := make([]int, 0)
	maxi := 0
	mini := 0
	fmt.Printf("Combien de notes voulez-vous saisir ?")
	fmt.Scan(&nombre)
	for cpt <= nombre {
		fmt.Println("Note", cpt, ":")
		fmt.Scan(&note)
		if note >= 0 && note <= 20 {
			lst = append(lst, note)
			somme += note
		} else {
			fmt.Println("note invalide")
			cpt--
		}
		cpt++
	}
	moyenne = calculerMoyenne(somme, nombre)
	maxi = trouverMaximum(lst)
	mini = trouverMinimum(lst)
	afficherResultat(moyenne, maxi, mini)

}
