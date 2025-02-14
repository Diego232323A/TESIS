# Update User Service 🛠️

Microservicio que permite actualizar un usuario.

## 🚀 Instalación y Ejecución

### 1️⃣ Configurar Variables de Entorno
Crea un archivo **`.env`** en la raíz:

```env
PORT=
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
JWT_SECRET=
```

### 2️⃣ Instalar Dependencias
```bash
npm install
```

### 3️⃣ Ejecutar
```bash
npm run start
```

## 📖 API Endpoints
### 📝 Documentación con Swagger
```bash
http://localhost:4004/api-docs
```

### 🛠️ Endpoints
``PUT	/api/user/:username	Actualizar usuario por username (requiere JWT)
``