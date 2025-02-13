# Microservice: Creating Reviews

This microservice allows the creation of reviews for tourist places, validating that the user and the place already exist.

## 🚀 Technologies used
- **Lenguage**: Go
- **Data Base**: PostgreSQL
- **Framework**: Gorilla Mux
- **Swagger**: Para documentación de API
- **Docker**: Para contenedorización

## 📌 Installation and execution

### 1️⃣ Clone the repository
```bash
git clone https://github.com/usuario/create-review.git
cd create-review
```

### 2️⃣ Setting environment variables
```bash
Create a .env file in the root of the project:
DB_HOST=
DB_NAME=
DB_USER=
DB_PASSWORD=
DB_PORT=
VALIDATION_SERVICE_URL=http://localhost:8080
```

### 📖 Endpoints (Swagger)
```bash
Documentation API: http://localhost:5005/swagger/index.html
```