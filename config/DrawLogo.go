package config

import (
	"io/ioutil"
	"my-api/pkg"
)

func DrawLogo() {
	content, err := ioutil.ReadFile("config/asset/logo.txt")
	if err != nil {
		pkg.DisplayContext("Error while reading file: config/asset/logo.txt", pkg.Error, err, true)
	}
	pkg.DisplayContext(string(content), pkg.Default)
}
