# Utilizar la imagen oficial de Go
FROM golang:1.21

# Establecer el directorio de trabajo
WORKDIR /app

# Instalar git si no está instalado en la imagen base
RUN apt-get update && apt-get install -y git

# Copiar los archivos go.mod y go.sum para aprovechar la caché de Docker
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar todos los archivos del proyecto al contenedor
COPY . .

# Descargar dependencias y limpiar módulos
RUN go mod tidy

# Compilar la aplicación
RUN go build -o consult-review .

# Exponer el puerto del servicio
EXPOSE 5008

# Comando para ejecutar la aplicación
CMD ["/app/consult-review"]
