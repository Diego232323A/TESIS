# 🛠️ Get User Service - Obtener Usuario por Username

Microservicio que permite obtener la información de un usuario por su `username`.

## 🚀 Instalación y Ejecución

### 1️⃣ Configurar Variables de Entorno
Crea un archivo **`.env`** en la raíz:

```env
PORT=4003
DB_HOST=postgres
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=users_db
DB_PORT=5432
JWT_SECRET=supersecreto

2️⃣ Instalar Dependencias
``npm install

3️⃣ Ejecutar en Modo Desarrollo
``npm run dev

📖 API Endpoints
📝 Documentación con Swagger
``http://localhost:4003/api-docs

🛠️ Endpoints
``GET	/api/user/:username	Obtener usuario por username (requiere JWT)