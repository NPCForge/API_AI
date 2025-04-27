package config

import (
	"io/ioutil"
	"my-api/pkg"
)

func DrawLogo() {
	content, err := ioutil.ReadFile("config/asset/logo.txt")
	if err != nil {
		pkg.DisplayContext("Erreur lors de la lecture du fichier", pkg.Error, err, true)
	}
	pkg.DisplayContext(string(content), pkg.Default)
}
