package utils

import (
	"bufio"
	"fmt"
	"my-api/pkg"
	"os"
	"strings"
)

func status() {
	store := pkg.GetPopulation() // retourne un map[string]string

	println("------------ Status ------------------------------------------------------------")
	for key, value := range store {
		fmt.Printf("Clé : %s, Valeur : %s\n", key, value)
	}
	fmt.Printf("------------ %d actifs ------------------------------------------------------------\n", len(store))
}

func Commande() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		switch input {
		case "stop":
			fmt.Println("Arrêt du serveur...")
			os.Exit(0)
		case "status":
			// fmt.Println("Le serveur est actif.")
			status()
		default:
			fmt.Println("Commande inconnue :", input)
		}
	}
}
