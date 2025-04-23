package utils

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"my-api/config"
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
	color.Green("script list\t: Affiche les test unitaire disponible.")
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

func scriptList() {
	entries, err := os.ReadDir(config.GetEnvVariable("UNITTEST_FOLDER"))
	if err != nil {
		pkg.DisplayContext("Error:", pkg.Error, err, true)
	}

	for _, e := range entries {
		pkg.DisplayContext(e.Name(), pkg.File)
	}
}

func scriptRun(file string) {
	scriptDir := config.GetEnvVariable("UNITTEST_FOLDER")

	entries, err := os.ReadDir(scriptDir)
	if err != nil {
		pkg.DisplayContext("Error:", pkg.Error, err, true)
		return
	}

	found := false
	for _, entry := range entries {
		if strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())) == file {
			found = true
		}
	}

	if !found {
		pkg.DisplayContext("Unknown script: "+file, pkg.Error)
		return
	}

	pkg.DisplayContext("Running script: "+file, pkg.Update)

	scriptPath := filepath.Join(scriptDir, file+".py")
	cmd := exec.Command("python", scriptPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		pkg.DisplayContext("Execution error:", pkg.Error, err)
	}

	pkg.DisplayContext("Output:\n"+string(output), pkg.Update)
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

		if strings.HasPrefix(input, "script run ") {
			name := strings.TrimPrefix(input, "script run ")
			name = strings.TrimSpace(name)
			scriptRun(name)
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

		case "script list":
			scriptList()

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
