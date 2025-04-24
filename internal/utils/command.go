package utils

import (
	"bufio"
	"os"
	"strings"

	"my-api/internal/services"
	"my-api/pkg"

	"github.com/fatih/color"
)

func status() {
	store := pkg.GetPopulation() // map[string]string

	color.Cyan("------------ 📊 Status ------------------------------------------------------------")
	for key, value := range store {
		color.Green("🔑 Clé : %s → 📦 Valeur : %s", key, value)
	}
	color.Cyan("------------ 🟢 %d actifs ------------------------------------------------------------\n", len(store))
}

func help() {
	color.Cyan("------------ ⌨️ Commandes ------------------------------------------------------------")
	color.Green("status\t: Retourne le nombre de personnes connectées.")
	color.Green("reset\t: Supprime tous les utilisateurs de la BDD et réinitialise le statut.")
	color.Green("stop\t: Coupe l'API.")
	color.Green("help\t: Affiche les informations sur les différentes commandes.")
	color.Cyan("-------------------------------------------------------------------------------------\n")
}

func reset() {
	rowsAffected, err := services.DropAllUser()
	if err != nil {
		color.Red("❌ Erreur lors de la requête SQL : %s", err)
		return
	}
	color.Cyan("💥 %d ligne(s) supprimée(s)", rowsAffected)
	pkg.ClearTokenStore()
	color.Cyan("💥 Tokenstore vidé.")
}

func Commande() {
	reader := bufio.NewReader(os.Stdin)

	color.Magenta("🧠 Console interactive prête. Tape une commande (help, stop, ...)\n")

	for {
		color.White("⤷ Entrez une commande : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.HasPrefix(input, "new route ") {
			name := strings.TrimPrefix(input, "new route ")
			name = strings.TrimSpace(name)
			GenerateNewRoute(name)
			continue
		}

		switch strings.ToLower(input) {
		case "stop":
			color.Red("⛔ Arrêt du serveur...")
			os.Exit(0)

		case "quit":
			color.Red("⛔ Arrêt du serveur...")
			os.Exit(0)

		case "exit":
			color.Red("⛔ Arrêt du serveur...")
			os.Exit(0)

		case "status":
			status()

		case "reset":
			reset()

		case "help":
			help()

		default:
			color.Yellow("❓ Commande inconnue : %s", input)
		}
	}
}
