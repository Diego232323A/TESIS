# 🌍 GeoService (Go + PostGIS)
Microservicio para gestionar geolocalización de lugares turísticos.

## 🚀 Tecnologías utilizadas
- **Lenguaje**: Go
- **Base de datos**: PostgreSQL
- **Swagger**: Para documentación de API
- **Docker**: Para contenedorización

## 🚀 Instalación
```bash
go mod tidy
go run main.go


## 📡 Endpoints
```bash
POST /places → Agrega un lugar.
GET /places/nearby?lat=...&lng=... → Obtiene lugares cercanos.

## 🛠 Swagger
```bash
Disponible en http://localhost:5012/swagger/index.html.