# Microservicio de Consulta de Lugares Turísticos

## Descripción
Este microservicio permite buscar lugares turísticos con filtros por nombre, ubicación y categoría.

## Instalación
1. Clonar el repositorio
2. Instalar dependencias: `pip install -r requirements.txt`
3. Configurar MongoDB en `settings.py`
4. Ejecutar el servidor: `uvicorn src.main:app --host 0.0.0.0 --port 5004 --reload`

## Endpoints
- `GET /api/places` → Permite buscar lugares turísticos aplicando filtros.

## Ejemplo de uso en Postman
### **Consulta con filtro de nombre**

## Buscar por nombre
    ``sh
    GET http://localhost:5004/api/places?name=Mitad

## Filtrar por ubicación
    ``sh
    GET http://localhost:5004/api/places?location=Quito

## Filtrar por categoría
    ``sh
    GET http://localhost:5004/api/places?category=Cultural