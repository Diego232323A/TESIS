# Imagen base
FROM python:3.10

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos al contenedor
COPY . /app

# Instalar dependencias
RUN pip install --no-cache-dir -r requirements.txt

# Exponer el puerto
EXPOSE 5003

# Comando para ejecutar la aplicación
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "5003", "--reload"]
