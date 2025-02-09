# Mi API de Búsqueda

Esta es una API para buscar en MongoDB y PostgreSQL.

## Requisitos

- Go 1.16 o superior
- Docker
- MongoDB
- PostgreSQL

## Instalación

1. Clona el repositorio:

```sh
git clone https://github.com/tu-usuario/mi-api-busqueda.git
cd mi-api-busqueda

2. Crea un archivo .env en la raíz del proyecto con el siguiente contenido:
```sh
MONGO_URI=
POSTGRES_HOST=
POSTGRES_DB=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_SSLMODE=disable


3. Construye y ejecuta el contenedor Docker:
```sh
docker build -t mi-api-busqueda .
docker run -p 8080:8080 mi-api-busqueda

4. La aplicación estará disponible en http://localhost:8080

5. La documentación de Swagger estará disponible en http://localhost:8080/swagger/index.html

## Enpoints

# POST /buscar
Busca en MongoDB y PostgreSQL.

{
    "valor_mongo": "valor_a_buscar_en_mongo",
    "valor_postgres": "valor_a_buscar_en_postgres"
}

