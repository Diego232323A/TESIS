# Usa la imagen oficial de Node.js
FROM node:16

# Crear un directorio de trabajo en el contenedor
WORKDIR /usr/src/app

# Copiar los archivos package.json y package-lock.json
COPY package*.json ./

# Instalar las dependencias
RUN npm install

# Copiar el resto del código fuente
COPY . .

# Exponer el puerto 8080
EXPOSE 8080

# Comando para iniciar el servidor
CMD ["npm", "start"]
