# 🛠️ Auth Service - Microservicio de Autenticación

Microservicio que maneja la autenticación de usuarios con JWT.

## 🚀 Instalación y Ejecución

### 1️⃣ Configurar Variables de Entorno
Crea un archivo **`.env`** en la raíz:

```env
PORT=4002
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
``http://localhost:4002/api-docs

🛠️ Endpoints
``POST	/auth/login	Iniciar sesión y obtener JWT
