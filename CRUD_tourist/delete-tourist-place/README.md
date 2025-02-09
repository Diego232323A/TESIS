# Microservicio - Eliminar Lugar Turístico

Este microservicio permite eliminar lugares turísticos utilizando FastAPI y MongoDB.

## Instalación y ejecución

1. Clonar el repositorio
   ```sh
   git clone https://github.com/tu-repo/delete-tourist-place.git
   cd delete-tourist-place/src

2. Instalar dependencias
    ```sh
    pip install -r requirements.txt

3. Ejecutar el servicio
    ```sh
    uvicorn main:app --host 0.0.0.0 --port 5003 --reload

4. Acceder a la API
    ```sh
    API: http://localhost:5003/api/place
    Swagger: http://localhost:5003/docs
