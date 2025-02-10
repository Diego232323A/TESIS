# Utilizar la imagen oficial de Go
FROM golang:1.21

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar todos los archivos del proyecto al contenedor
COPY . .

# Descargar dependencias
RUN go mod tidy

# Compilar la aplicación
RUN go build -o delete-review .

# Exponer el puerto del servicio
EXPOSE 5007

# Comando para ejecutar la aplicación
CMD ["/app/delete-review"]
