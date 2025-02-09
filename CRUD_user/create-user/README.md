# 🛠️ Crear Usuario - Microservicio

Microservicio REST para registrar usuarios en PostgreSQL.

## 🚀 Tecnologías
- **Node.js** con Express.js
- **PostgreSQL** como base de datos
- **Joi** para validaciones
- **Swagger** para documentación
- **Docker y Docker Compose**

## 📌 Instalación y Ejecución

### 1️⃣ Configurar Variables de Entorno
Crea un archivo **`.env`** en la raíz:

```env
PORT=4001
DB_HOST=postgres
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=users_db
DB_PORT=5432
CORS_ORIGIN=*

2️⃣ Instalar Dependencias
``npm install

3️⃣ Ejecutar en Modo Desarrollo
``npm run dev

📖 API Endpoints
📝 Documentación con Swagger
``http://localhost:4001/api-docs

🛠️ Endpoints
``POST	/users	Crear un nuevo usuario
