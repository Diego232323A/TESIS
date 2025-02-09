# Imagen base
FROM node:18

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar archivos y dependencias
COPY package.json package-lock.json ./
RUN npm install

# Copiar el resto del código
COPY . .

# Exponer el puerto
EXPOSE 4001

# Comando de inicio
CMD ["npm", "start"]
