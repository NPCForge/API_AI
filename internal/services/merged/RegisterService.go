package service

import "my-api/pkg"

func RegisterService(password string, identifiant string) (string, error) {
	pkg.DisplayContext("RegisterService", pkg.Debug)
	// creation de l'utilisateur
	// crée une clé api
	// retourne la clé
	return "", nil
}

// Http
// func RegisterService(entity httpModels.RegisterRequest) (string, error)

// Websocket
// func RegisterServiceWebSocket(
// 	conn *websocket.Conn,
// 	msg websocketModels.RegisterRequest,
// 	sendResponse types.SendResponseFunc,
// 	sendError types.SendErrorFunc,
// )

// Http
// func RegisterService(entity httpModels.RegisterRequest) (string, error) {
// 	id, err := SaveInDatabase(entity)

// 	if err != nil {
// 		return "", errors.New("error saving in database")
// 	}

// 	pass, err := utils.GenerateJWT(strconv.FormatInt(id, 10))

// 	if err != nil {
// 		return "", errors.New("error generating JWT")
// 	}

// 	pkg.SetToken(strconv.FormatInt(id, 10), pass)

// 	return pass, nil
// }

// Websocket version
// func RegisterServiceWebSocket(
// 	conn *websocket.Conn,
// 	msg websocketModels.RegisterRequest,
// 	sendResponse types.SendResponseFunc,
// 	sendError types.SendErrorFunc,
// ) {
// 	initialRoute := "Register"

// 	id, err := SaveInDatabase(msg)
// 	if err != nil {
// 		color.Red("❌ Failed to save in database: %v", err)
// 		sendError(conn, initialRoute, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	stringId := strconv.FormatInt(id, 10)

// 	pass, err := utils.GenerateJWT(stringId)
// 	if err != nil {
// 		color.Red("❌ Failed to generate JWT: %v", err)
// 		sendError(conn, initialRoute, map[string]interface{}{
// 			"message": "Unable to generate Token",
// 		})
// 		return
// 	}

// 	pkg.SetToken(stringId, pass)

// 	color.Green("✅ Token generated and stored for user: %s", stringId)
// 	sendResponse(conn, initialRoute, map[string]interface{}{
// 		"token": pass,
// 	})
// }
