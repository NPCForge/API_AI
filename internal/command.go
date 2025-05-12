package internal

import (
	"bufio"
	"my-api/internal/services"
	"my-api/internal/utils"
	"os"
	"strings"

	"my-api/pkg"

	"github.com/fatih/color"
)

// status displays the number of currently connected users.
func status() {
	store := pkg.GetPopulation() // map[string]string

	color.Cyan("------------ 📊 Status ------------------------------------------------------------")
	for key, value := range store {
		color.Green("🔑 Key: %s → 📦 Value: %s", key, value)
	}
	color.Cyan("------------ 🟢 %d active users ---------------------------------------------------\n", len(store))
}

// help displays all available commands.
func help() {
	color.Cyan("------------ ⌨️ Commands ------------------------------------------------------------")
	color.Green("status\t: Returns the number of connected users.")
	color.Green("reset\t: Deletes all users from the database and resets the status.")
	color.Green("stop\t: Stops the API.")
	color.Green("resetdiscussions\t: Deletes all discussions for all users.")
	color.Green("help\t: Displays information about available commands.")
	color.Cyan("-------------------------------------------------------------------------------------\n")
}

// reset deletes all users from the database and clears the token store.
func reset() {
	rowsAffected, err := services.DropAllUser()
	if err != nil {
		color.Red("❌ Error during SQL request: %s", err)
		return
	}
	color.Cyan("💥 %d row(s) deleted", rowsAffected)
	pkg.ClearTokenStore()
	color.Cyan("💥 Token store cleared.")
}

// resetDiscussions deletes all existing discussions.
func resetDiscussions() {
	err := services.DropDiscussions()
	if err != nil {
		pkg.DisplayContext("Cannot reset discussions:", pkg.Error, err)
		return
	}
	pkg.DisplayContext("Discussions successfully deleted!", pkg.Update)
}

// Commande launches the interactive console for server administration.
func Commande() {
	reader := bufio.NewReader(os.Stdin)

	color.Magenta("🧠 Interactive console ready. Type a command (help, stop, ...)\n")

	for {
		color.White("⤷ Enter a command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.HasPrefix(input, "new route ") {
			name := strings.TrimPrefix(input, "new route ")
			name = strings.TrimSpace(name)
			utils.GenerateNewRoute(name)
			continue
		}

		switch strings.ToLower(input) {
		case "stop", "quit", "exit":
			color.Red("⛔ Shutting down the server...")
			os.Exit(0)

		case "status":
			status()

		case "reset":
			reset()

		case "resetdiscussions":
			resetDiscussions()

		case "help":
			help()

		default:
			color.Yellow("❓ Unknown command: %s", input)
		}
	}
}
