# Usa la imagen base de Node.js
FROM node:18

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el package.json y package-lock.json
COPY package.json ./

# Instala las dependencias
RUN npm install

# Copia todo el código fuente al contenedor
COPY . .

# Expone el puerto 5009
EXPOSE 5009

# Comando de inicio
CMD ["npm", "start"]
