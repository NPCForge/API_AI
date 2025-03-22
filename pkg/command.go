package pkg

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func status() {
	store := GetPopulation()

	color.Cyan("------------ 📊 Status ------------------------------------------------------------")
	for key, value := range store {
		color.Green("🔑 Clé : %s → 📦 Valeur : %s", key, value)
	}
	color.Cyan("------------ 🟢 %d actifs ------------------------------------------------------------\n", len(store))
}

func Commande() {
	reader := bufio.NewReader(os.Stdin)

	color.Magenta("🧠 Console interactive prête. Tape une commande (status, stop, ...)\n")

	for {
		color.White("⤷ Entrez une commande : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		switch input {
		case "stop":
			color.Red("⛔ Arrêt du serveur...")
			os.Exit(0)

		case "status":
			status()

		default:
			color.Yellow("❓ Commande inconnue : %s", input)
		}
	}
}
