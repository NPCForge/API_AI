package config

import (
	"fmt"
	"io/ioutil"
	"log"
)

func DrawLogo() {
	content, err := ioutil.ReadFile("config/asset/logo.txt")
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}

	fmt.Println(string(content))
}
