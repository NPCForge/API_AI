# **Go API Project**

## **Description**
This project is a Go-based API that interacts with a PostgreSQL database. It provides functionality to manage entities and uses Docker to simplify deployment.

---

## **Prerequisites**
1. **Docker** and **Docker Compose** should be installed.
   - [Install Docker](https://docs.docker.com/get-docker/)
   - [Install Docker Compose](https://docs.docker.com/compose/install/)

2. A working PostgreSQL instance (local or remote).
   - You can use a local PostgreSQL server, a hosted instance, or set it up separately in Docker if needed.

---

## **Configuration**

### 1. Environment Variables
An example `.env` file is provided as `.env.example`. Copy it and adjust the values as needed:

```bash
cp .env.example .env
```

**Example `.env.example` file:**
```env
DB_HOST=localhost       # PostgreSQL host address
DB_PORT=5432            # PostgreSQL port
DB_USER=API             # PostgreSQL username
DB_PASSWORD=password    # PostgreSQL password
DB_NAME=api_db          # PostgreSQL database name

APP_PORT=8080           # API listening port
```

### 2. PostgreSQL Setup
#### Create the Database
1. Create a PostgreSQL database with the information in `.env`:
   ```sql
   CREATE DATABASE api_db;
   CREATE USER API WITH PASSWORD 'password';
   GRANT ALL PRIVILEGES ON DATABASE api_db TO API;
   ```

2. Apply the migration to create the required table:
   The `migrations/CREATE TABLE entity.sql` file contains the table schema. Run the following command to apply it:
   ```bash
   psql -U API -d api_db -f migrations/CREATE\ TABLE\ entity.sql
   ```

   **Table structure:**
   - Table name: `entity`
   - Columns:
     - `id` (primary key)
     - `nom` (name of the entity)
     - `token` (unique token for the entity)
     - `prompt` (associated text)
     - `created` (creation date)

---

## **Running the Project**

### 1. Build and Run the API with Docker
1. Build and start the container:
   ```bash
   docker-compose up --build
   ```

2. The API will be available at:
   [http://localhost:8080](http://localhost:8080)

---

## **Available Endpoints**
### **Entity**
- **Create an Entity**  
  **POST** `/entity`
  ```json
  {
    "name": "New Entity",
    "token": "unique-token",
    "prompt": "This is a description."
  }
  ```

- **Retrieve an Entity**  
  **GET** `/entity/:id`

- **Delete an Entity**  
  **DELETE** `/entity/:id`

---

## **Stopping the Project**
Stop the Docker containers:
```bash
docker-compose down
```

---

## **Troubleshooting**

### Common Issues
1. **PostgreSQL Connection Error**:
   - Ensure the database is reachable at the host and port specified in `.env`.
   - If using a remote database, make sure its firewall allows connections from the server running Docker.

2. **Missing `.env` File**:
   - Ensure the file exists and is correctly configured. You can use `.env.example` as a template.

---

## **Project Structure**
```plaintext
.
├── cmd/
│   └── app/         # Main entry point
├── config/          # Database and environment variable configuration
├── internal/
│   ├── handlers/    # HTTP request handlers
│   ├── models/      # Data models
│   └── services/    # Business logic
├── migrations/      # SQL migration files
├── pkg/             # Utility functions (e.g., JWT handling)
├── Dockerfile       # Instructions to build the Docker image
├── docker-compose.yml # Docker Compose configuration
├── go.mod           # Go project dependencies
├── .env.example     # Example environment variables file
├── .env       # Local environment variables file
└── README.md        # Project documentation
```

---

## **Contributions**
Contributions are welcome! To propose changes:
1. Fork this repository.
2. Create a new branch (`git checkout -b feature/improvement`).
3. Commit your changes (`git commit -m 'Add a feature'`).
4. Submit a Pull Request.

---

If you have any issues or need further explanations, feel free to ask! 😊

Discord : cogal
