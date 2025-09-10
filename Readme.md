# API_AI

API open-source en Go permettant de piloter des entités dans un jeu de type Loups-Garous. Les entités disposent d'un rôle et d'un prompt et peuvent discuter ou voter grâce au modèle d'OpenAI. L'API expose des routes HTTP ainsi qu'un WebSocket.

## Fonctionnalités
- Inscription et connexion d'utilisateurs
- Création, suppression et listing d'entités
- Génération de décisions (discussion, vote...) via ChatGPT
- Échange temps réel par WebSocket (`/ws`)

## Installation

### Prérequis
- Go ≥ 1.23
- Docker & Docker Compose *(facultatif mais recommandé)*
- PostgreSQL 15
- Clé API OpenAI (`CHATGPT_TOKEN`)

### Cloner le dépôt
```bash
git clone <URL_DU_DEPOT>
cd API_AI
```

### Configuration
Copier le fichier d'exemple et renseigner vos valeurs:
```bash
cp .env.exemple .env
```
Variables importantes:
- `API_KEY_REGISTER` : clé partagée pour créer un utilisateur
- `JWT_SECRET_KEY` : secret pour signer les tokens
- `CHATGPT_TOKEN` : clé OpenAI
- `POSTGRES_*` : accès à la base de données

### Démarrage avec Docker
```bash
docker-compose up --build
```
La base PostgreSQL est initialisée depuis `migrations/` et l'API écoute sur `http://localhost:3000`.

### Démarrage sans Docker
1. Démarrer un serveur PostgreSQL et appliquer les migrations du dossier `migrations/`
2. Installer les dépendances Go:
```bash
go mod download
```
3. Lancer l'API:
```bash
go run ./cmd/app
```

## Utilisation

### Routes HTTP principales
- `POST /Register` – créer un utilisateur
- `POST /Connect` – obtenir un token
- `POST /Disconnect` – invalider un token
- `POST /CreateEntity` – créer une entité
- `POST /RemoveEntity` – supprimer une entité
- `GET /GetEntities` – lister les entités
- `POST /MakeDecision` – demander une action à une entité
- `GET /Status` – vérifier l'authentification
- `GET /health` – vérifier l'état du serveur

### WebSocket
Connexion sur `ws://localhost:3000/ws`. Envoyer des messages JSON contenant un champ `Action` pour appeler les mêmes opérations (Register, Connect, MakeDecision, etc.).

## Tests
```bash
go test $(go list ./... | grep -v script)
```

## Licence et contributions
Projet ouvert aux contributions. Les propositions de correctifs ou de fonctionnalités sont les bienvenues via Pull Request.
