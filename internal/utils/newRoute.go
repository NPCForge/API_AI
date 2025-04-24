package utils

import (
	"fmt"
	"io"
	"my-api/config"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	// Ouvrir le fichier source
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("erreur en ouvrant le fichier source : %w", err)
	}
	defer sourceFile.Close()

	// Créer le fichier de destination
	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("erreur en créant le fichier destination : %w", err)
	}
	defer destFile.Close()

	// Copier le contenu
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("erreur lors de la copie : %w", err)
	}

	// S'assurer que tout est bien écrit
	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("erreur lors du sync : %w", err)
	}

	return nil
}

func copyWithRename(srcBase, srcName, dstBase, dstName string) error {
	src := filepath.Join(srcBase, srcName)
	dst := filepath.Join(dstBase, dstName)
	err := copyFile(src, dst)
	if err != nil {
		return fmt.Errorf("échec de la copie de %s vers %s : %w", src, dst, err)
	}
	return nil
}

func GenerateNewRoute(name string) {
	basePath := config.GetEnvVariable("PATH_EXEMPLE")

	files := []struct {
		srcFile     string
		envDestPath string
		suffix      string
	}{
		{"handlerWebSocket.go", "PATH_WS_HANDLER", "Handler.go"},
		{"handlerHttp.go", "PATH_HTTP_HANDLER", "Handler.go"},
		{"Model.go", "PATH_MODEL", "Model.go"},
		{"Service.go", "PATH_SERVICE", "Service.go"},
	}

	for _, f := range files {
		destPath := config.GetEnvVariable(f.envDestPath)
		destFile := name + f.suffix
		if err := copyWithRename(basePath, f.srcFile, destPath, destFile); err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Printf("✔ %s copié dans %s\n", f.srcFile, filepath.Join(destPath, destFile))
		}
	}
}
