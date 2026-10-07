package main

import ( 
		"os"
		"fmt"
)
func main () {
	if len(os.Args) != 3 {
		fmt.Println("Usage de la fonction : go run . <fichier_depart> <nom_fichier_arrivee>")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Erreur lors de la lecture :", err) 
		return
	}
	
	cleanText := ProcessText(string(data))

	err = os.WriteFile(output, []byte(cleanText), 0644)
	if err != nil {
		fmt.Println("Erreur lros de l'écriture du fichier :", err)
	}
}