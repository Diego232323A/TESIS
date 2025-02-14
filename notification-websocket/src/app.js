const express = require("express");
const http = require("http");
const mongoose = require("mongoose");
const WebSocket = require("ws");
const { Pool } = require("pg");
const { initWebSocket, saveNotification } = require("./websocket");

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

// Middleware para manejar JSON en el body de las solicitudes
app.use(express.json());

// Conectar a MongoDB
mongoose.connect("mongodb://127.0.0.1:27017/notifications_db", {
    useNewUrlParser: true,
    useUnifiedTopology: true,
});

mongoose.connection.on("connected", () => {
    console.log("Conectado a MongoDB");
});

// Inicializar WebSocket
initWebSocket(wss);

// Configurar la conexión a PostgreSQL
const pool = new Pool({
    user: 'postgres',         // Cambia esto por tu usuario de PostgreSQL
    host: 'localhost',          // Cambia esto por la dirección de tu servidor PostgreSQL
    database: 'users_db', // Cambia esto por tu nombre de base de datos
    password: '12345',  // Cambia esto por tu contraseña de PostgreSQL
    port: 5432,                 // Cambia esto si tu servidor PostgreSQL está en otro puerto
});

app.post("/send-notification", async (req, res) => {
    const { message, userId } = req.body;
    if (!message || !userId) {
        return res.status(400).send("Faltan parámetros.");
    }

    try {
        // Obtener información del usuario desde PostgreSQL
        const userQuery = await pool.query('SELECT * FROM users WHERE id = $1', [userId]);
        const user = userQuery.rows[0];

        if (!user) {
            return res.status(404).send("Usuario no encontrado.");
        }

        // Guardar la notificación en MongoDB y enviarla a los clientes WebSocket
        await saveNotification(message, user);
        res.status(200).send("Notificación enviada.");
    } catch (error) {
        console.error("Error en /send-notification:", error);
        res.status(500).send(`Error al guardar la notificación: ${error.message}`);
    }
});


// Iniciar servidor
server.listen(8081, () => {
    console.log("Servidor escuchando en http://localhost:8081");
});
