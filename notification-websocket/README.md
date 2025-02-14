# Microservicio WebSocket - Notificaciones en Tiempo Real

Este microservicio proporciona un sistema de notificaciones en tiempo real utilizando WebSocket y MongoDB. Recibe notificaciones y las envía a los clientes conectados.

## Tecnologías utilizadas:
- **Node.js**
- **WebSocket (ws)**
- **MongoDB** (para almacenar las notificaciones)
- **Express** (para crear el servidor HTTP)

## Instalación

1. Clona el repositorio:
   ```bash
   git clone <URL-del-repositorio>
   cd websocket-service

2. Configurar Variables de Entorno:

Crea un archivo **`.env`** en la raíz:

```env
PORT=4003
DB_HOST=postgres
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=users_db
DB_PORT=5432

3. Install Dependences 
``npm install

4. Run
``npm run start