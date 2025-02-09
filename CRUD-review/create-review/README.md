# Microservicio: Creación de Reseñas

Este microservicio permite la creación de reseñas para lugares turísticos, validando que el usuario y el lugar existan previamente.

## 🚀 Tecnologías utilizadas
- **Lenguaje**: Go
- **Base de datos**: PostgreSQL
- **Framework**: Gorilla Mux
- **Swagger**: Para documentación de API
- **Docker**: Para contenedorización

## 📌 Instalación y ejecución

### 1️⃣ Clonar el repositorio
```bash
git clone https://github.com/usuario/create-review.git
cd create-review

### 2️⃣ Configurar las variables de entorno
```bash
Crear un archivo .env en la raíz del proyecto:
DB_HOST=localhost
DB_NAME=reviews_db
DB_USER=postgres
DB_PASSWORD=12345
DB_PORT=5432
VALIDATION_SERVICE_URL=http://localhost:8080

📖 Endpoints (Swagger)
```bash
Documentación API: http://localhost:5005/swagger/index.html