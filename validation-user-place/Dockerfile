# Usa una imagen base oficial de Go
FROM golang:1.16

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente del proyecto al contenedor
COPY . .

# Expone el puerto en el que la aplicación estará escuchando
EXPOSE 8080

# Compila la aplicación
RUN go build -o main .

# Comando por defecto al iniciar el contenedor
CMD ["./main"]
