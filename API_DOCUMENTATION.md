# Documentation API - HTTP & WebSocket

## 📋 Table des matières
- [Informations générales](#informations-générales)
- [Routes HTTP](#routes-http)
- [Routes WebSocket](#routes-websocket)
- [Modèles de données](#modèles-de-données)
- [Codes de statut](#codes-de-statut)

---

## 🌐 Informations générales

### Configuration du serveur
- **URL de base**: `http://0.0.0.0:3000`
- **CORS**: Autorisé depuis `http://localhost:8001`
- **Méthodes CORS autorisées**: `GET`, `POST`, `OPTIONS`
- **Headers autorisés**: `Content-Type`, `Authorization`

### Authentification
Les routes protégées nécessitent un token JWT dans le header `Authorization`.

---

## 🔗 Routes HTTP

### Routes publiques

#### `POST /Connect`
Connexion d'un utilisateur existant.

**Body**:
```json
{
  "identifier": "string",
  "password": "string"
}
```

**Réponse**:
```json
{
  "message": "Successfully connected",
  "status": 200,
  "id": "string",
  "TmpToken": "string"
}
```

**Codes de statut**:
- `200` : Connexion réussie
- `401` : Identifiants invalides
- `400` : Requête invalide
- `405` : Méthode non autorisée

---

#### `POST /Register`
Inscription d'un nouvel utilisateur (nécessite une clé API).

**Body**:
```json
{
  "API_KEY": "string",
  "identifier": "string",
  "password": "string",
  "game_prompt": "string"
}
```

**Réponse**:
```json
{
  "message": "User created",
  "status": 201,
  "token": "string",
  "id": "string"
}
```

**Codes de statut**:
- `201` : Utilisateur créé
- `204` : Données manquantes
- `401` : Clé API invalide
- `400` : Requête invalide
- `500` : Erreur serveur

---

### Routes protégées (authentification requise)

#### `POST /Disconnect`
Déconnexion d'un utilisateur.

**Headers**:
```
Authorization: Bearer <token>
```

**Body**:
```json
{
  "token": "string"
}
```

**Réponse**:
```json
{
  "message": "string",
  "status": 200
}
```

---

#### `POST /RemoveUser`
Suppression d'un utilisateur.

**Headers**:
```
Authorization: Bearer <token>
```

**Body**:
```json
{
  "token": "string",
  "username": "string"
}
```

**Réponse**:
```json
{
  "message": "string",
  "status": 200
}
```

---

#### `POST /MakeDecision`
Permet à une entité de prendre une décision basée sur un message.

**Headers**:
```
Authorization: Bearer <token>
```

**Body**:
```json
{
  "message": "string",
  "checksum": "string"
}
```

**Réponse**:
```json
{
  "Status": "200",
  // Données additionnelles retournées par le service
}
```

**Codes de statut**:
- `200` : Décision prise avec succès
- `400` : Requête invalide ou champs manquants

---

#### `POST /CreateEntity`
Création d'une nouvelle entité.

**Headers**:
```
Authorization: Bearer <token>
```

**Body**:
```json
{
  "name": "string",
  "prompt": "string",
  "checksum": "string",
  "role": "string"
}
```

**Réponse**:
```json
{
  "id": "string",
  "message": "Entity created successfully",
  "checksum": "string",
  "status": "success"
}
```

**Codes de statut**:
- `200` : Entité créée
- `400` : Requête invalide
- `500` : Erreur lors de la création

---

#### `POST /RemoveEntity`
Suppression d'une entité.

**Headers**:
```
Authorization: Bearer <token>
```

**Body**:
```json
{
  "checksum": "string",
  "token": "string"
}
```

**Réponse**:
```json
{
  "message": "string",
  "status": 200
}
```

---

#### `GET /GetEntities`
Récupération de toutes les entités associées à un utilisateur.

**Headers**:
```
Authorization: Bearer <token>
```

**Réponse**:
```json
{
  "entity": [
    {
      // Données des entités
    }
  ],
  "status": "success"
}
```

**Codes de statut**:
- `200` : Succès
- `500` : Erreur serveur
- `405` : Méthode non autorisée

---

#### `GET /Status`
Vérification du statut d'authentification.

**Headers**:
```
Authorization: Bearer <token>
```

**Réponse**:
```json
{
  "message": "Authentified",
  "status": 200
}
```

---

### Route de santé

#### `GET /health`
Vérification de l'état du serveur.

**Réponse**:
```
OK
```

**Code de statut**: `200`

---

## 🔌 Routes WebSocket

### Connexion WebSocket
**URL**: `ws://0.0.0.0:3000/ws`

Toutes les communications WebSocket utilisent le format suivant :

```json
{
  "action": "NomDeLAction",
  // Autres paramètres selon l'action
}
```

---

### Actions disponibles

#### `Register` (non protégée)
Inscription d'un nouvel utilisateur via WebSocket.

**Message**:
```json
{
  "action": "Register",
  "API_KEY": "string",
  "identifier": "string",
  "password": "string",
  "game_prompt": "string"
}
```

---

#### `Connect` (non protégée)
Connexion d'un utilisateur via WebSocket.

**Message**:
```json
{
  "action": "Connect",
  "identifier": "string",
  "password": "string"
}
```

---

#### `Disconnect` (protégée)
Déconnexion d'un utilisateur.

**Message**:
```json
{
  "action": "Disconnect",
  "token": "string"
}
```

---

#### `MakeDecision` (protégée)
Prise de décision par une entité.

**Message**:
```json
{
  "action": "MakeDecision",
  "message": "string",
  "checksum": "string"
}
```

---

#### `RemoveUser` (protégée)
Suppression d'un utilisateur.

**Message**:
```json
{
  "action": "RemoveUser",
  "token": "string",
  "username": "string"
}
```

---

#### `CreateEntity` (protégée)
Création d'une entité.

**Message**:
```json
{
  "action": "CreateEntity",
  "name": "string",
  "prompt": "string",
  "checksum": "string",
  "role": "string"
}
```

---

#### `RemoveEntity` (protégée)
Suppression d'une entité.

**Message**:
```json
{
  "action": "RemoveEntity",
  "checksum": "string",
  "token": "string"
}
```

---

#### `Restart` (protégée)
Redémarrage d'une session ou d'un processus.

**Message**:
```json
{
  "action": "Restart"
}
```

---

#### `GetEntities` (protégée)
Récupération des entités.

**Message**:
```json
{
  "action": "GetEntities"
}
```

---

#### `Status` (protégée)
Vérification du statut.

**Message**:
```json
{
  "action": "Status"
}
```

---

## 📊 Modèles de données

### ConnectRequest
```json
{
  "identifier": "string",
  "password": "string"
}
```

### ConnectResponse
```json
{
  "message": "string",
  "status": number,
  "id": "string",
  "TmpToken": "string"
}
```

### RegisterRequest
```json
{
  "API_KEY": "string",
  "identifier": "string",
  "password": "string",
  "game_prompt": "string"
}
```

### RegisterResponse
```json
{
  "message": "string",
  "status": number,
  "token": "string",
  "id": "string"
}
```

### CreateEntityRequest
```json
{
  "name": "string",
  "prompt": "string",
  "checksum": "string",
  "role": "string"
}
```

### CreateEntityResponse
```json
{
  "id": "string",
  "message": "string",
  "checksum": "string",
  "status": "success" | "error"
}
```

### MakeDecisionRequest
```json
{
  "message": "string",
  "checksum": "string"
}
```

### GetEntitiesResponse
```json
{
  "entity": Array<Entity>,
  "status": "success" | "error"
}
```

---

## 📌 Codes de statut

### Codes de succès
- `200` : OK
- `201` : Created

### Codes d'erreur client
- `400` : Bad Request
- `401` : Unauthorized
- `404` : Not Found
- `405` : Method Not Allowed

### Codes d'erreur serveur
- `500` : Internal Server Error

---

## 🔐 Sécurité

### Authentification
Les routes protégées nécessitent un token JWT dans le header `Authorization`:
```
Authorization: Bearer <votre_token_jwt>
```

### Middleware
- **LoggingMiddleware** : Appliqué sur toutes les routes protégées
- **CORS** : Configuration pour permettre les requêtes depuis `localhost:8001`

---

## 📝 Notes importantes

1. **WebSocket vs HTTP** : Les actions WebSocket offrent les mêmes fonctionnalités que les routes HTTP mais via une connexion persistante.

2. **Protection des routes** : Les routes WebSocket marquées comme "protégées" nécessitent une authentification préalable.

3. **Clé API** : L'inscription nécessite une clé API (`API_KEY_REGISTER`) définie dans les variables d'environnement.

4. **Format des erreurs** : Les erreurs WebSocket sont envoyées avec le format :
   ```json
   {
     "error": true,
     "message": "string"
   }
   ```

5. **Checksum** : Utilisé comme identifiant unique pour les entités dans plusieurs opérations.

---

## 🚀 Démarrage rapide

### Exemple de connexion HTTP
```bash
curl -X POST http://0.0.0.0:3000/Connect \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "user@example.com",
    "password": "password123"
  }'
```

### Exemple de connexion WebSocket (JavaScript)
```javascript
const ws = new WebSocket('ws://0.0.0.0:3000/ws');

ws.onopen = () => {
  ws.send(JSON.stringify({
    action: 'Connect',
    identifier: 'user@example.com',
    password: 'password123'
  }));
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log(data);
};
```

---

**Dernière mise à jour** : Février 2026
