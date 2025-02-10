# Microservicio - Eliminar Lugar Turístico

Este microservicio permite actualizar lugares turísticos utilizando FastAPI y MongoDB.

## Instalación y ejecución

1. Clonar el repositorio
   ```sh
   git clone https://github.com/tu-repo/update-tourist-place.git
   cd update-tourist-place/src

2. Instalar dependencias
    ```sh
    pip install -r requirements.txt

3. Ejecutar el servicio
    ```sh
    uvicorn main:app --host 0.0.0.0 --port 5002 --reload

4. Acceder a la API
    ```sh
    API: http://localhost:5002/api/place
    Swagger: http://localhost:5002/docs
