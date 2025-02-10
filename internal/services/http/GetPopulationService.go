package httpServices

import "my-api/pkg"

func GetPopulation() map[string]string {
	// aller chercher dans le store les personne connecter et les retourner
	store := pkg.GetPopulation()
	return store
}
