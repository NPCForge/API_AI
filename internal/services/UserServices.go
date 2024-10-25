package services

import "my-api/pkg"

func RegisterService() string {
	pass, err := pkg.GenerateJWT("albert")
	if err != nil {
		return "error generating JWT"
	}
	return pass
}
