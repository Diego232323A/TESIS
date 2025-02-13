# Usar una imagen base de Golang
FROM golang:1.20

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Exponer el puerto donde corre el servicio
EXPOSE 5005

# Comando para ejecutar el microservicio
CMD ["./create-review"]
