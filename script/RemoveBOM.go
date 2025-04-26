package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	bom = []byte{0xEF, 0xBB, 0xBF}
)

func main() {
	root := ".." // Change ici si tu veux scanner ailleurs

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if bytes.HasPrefix(data, bom) {
			fmt.Println("🚀 BOM détecté et supprimé :", path)
			err := ioutil.WriteFile(path, data[len(bom):], 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erreur:", err)
	}
}
