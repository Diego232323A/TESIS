# 🗑️ Microservicio: Eliminar Reseña

Este microservicio permite **eliminar reseñas de lugares turísticos**, asegurando que **solo el propietario de la reseña pueda eliminarla**.

## 🚀 Tecnologías
- Go 1.21+
- PostgreSQL
- Gorilla Mux (Enrutamiento)
- Swagger (Documentación)
- CORS

## 📌 Instalación y Ejecución

### **1️⃣ Configurar la Base de Datos**
Este microservicio se conecta a PostgreSQL. Asegúrate de que la tabla `reviews` esté creada:

```sql
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    rating INT NOT NULL,
    comment TEXT NOT NULL
);


### **2️⃣ Clonar el Repositorio**

```sh
git clone https://github.com/tu-repo/delete-review.git
cd delete-review

### **3️⃣ Configurar Variables de Entorno**

#### ·Crea un archivo .env con las credenciales de la base de datos

### *4️⃣ Inicia el repositorio*

```sh
go mod init nombre-de-tu-prouecto    

### *5️⃣ Instala las dependecias*

```sh
go mod tidy

### Por ultimo ejecuta el codigo.

```sh
go run main.go
