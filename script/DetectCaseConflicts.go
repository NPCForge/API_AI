package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fileInfo struct {
	path    string
	pkgName string
}

func main() {
	root := ".." // tu peux changer si besoin

	filesMap := make(map[string]fileInfo)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// On ignore les dossiers
		if info.IsDir() {
			return nil
		}

		// On ne prend que les fichiers .go
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		// Lisons le package d�clar� dans le fichier
		packageName, err := readPackageName(path)
		if err != nil {
			fmt.Println("Erreur lecture package pour", path, ":", err)
			return nil // continue malgr� l'erreur
		}

		key := strings.ToLower(info.Name()) + ":" + packageName

		if previous, exists := filesMap[key]; exists {
			fmt.Println("??  Collision d�tect�e dans le m�me package !")
			fmt.Println("   -", previous.path)
			fmt.Println("   -", path)
		} else {
			filesMap[key] = fileInfo{path: path, pkgName: packageName}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erreur pendant le scan:", err)
	}
}

func readPackageName(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignore les lignes vides ou commentaire
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		if strings.HasPrefix(line, "package ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "package ")), nil
		}
	}

	return "", fmt.Errorf("aucun package trouv�")
}
