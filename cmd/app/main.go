package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Erreur de chargement du fichier .env.local")
	}

	// Route principale
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Message sended!")
		return c.SendString("Hello, World!")
	})

	// Démarrer le serveur
	log.Fatal(app.Listen(":3000"))
}
