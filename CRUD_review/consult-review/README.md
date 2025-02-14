# 📖 Microservicio: Consultar Reseñas

Este microservicio permite **consultar reseñas de lugares turísticos** en PostgreSQL.

## 🚀 Tecnologías
- Go 1.21+
- PostgreSQL
- Gorilla Mux
- Swagger
- CORS

## 📌 Instalación y Ejecución

### **1️⃣ Configurar la Base de Datos**
```sql
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    place_id VARCHAR(255) NOT NULL,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comment TEXT NOT NULL
);
```

### **2️⃣ Configurar Variables de Entorno**

#### ·Create a .env file with the database credentials

### *3️⃣ Start the repository*
```sh
go mod init nombre-de-tu-prouecto    
```

### *4️⃣ Install dependences*
```sh
go mod tidy
```

### 5️⃣Run
```sh
go run main.go
```
