# 🛠️ Monitoring Microservice

This microservice is designed for monitoring using Grafana and Prometheus. It is built with Node.jsand uses environment variables defined in a .env file. The service is containerized using Docker.

## 📌 Prerequisites
-**Node.js**
-**Docker**
-**Prometheus**
-**Grafana**

## Setup
### 1️⃣ Configure Environment Variables
Create a .env file in the root of the project and add the necessary environment variables:
```bash
PORT=
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
```

### 2️⃣ Install Dependes
```bash
npm install
```

### 3️⃣ Run
```bash
npm start
```
 
### Access the Services
send alert
```bash
http://localhost:5009/alert
```