const WebSocket = require("ws");
const Notification = require("./models/notification");

const clients = new Set();

const initWebSocket = (wss) => {
    wss.on("connection", (ws) => {
        clients.add(ws);
        console.log("🔗 Cliente WebSocket conectado");

        // Enviar notificación a todos los clientes conectados
        ws.on("message", (message) => {
            console.log(`Mensaje recibido: ${message}`);
        });

        ws.on("close", () => {
            clients.delete(ws);
            console.log("❌ Cliente WebSocket desconectado");
        });
    });
};

const sendNotification = (notification) => {
    clients.forEach((client) => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(notification));
        }
    });
};

// Guardar notificación en MongoDB
const saveNotification = async (message, user) => {
    try {
        const newNotification = new Notification({
            message,
            user: user.id, // Utiliza el ID del usuario aquí
            timestamp: new Date(),
        });
        await newNotification.save();
        sendNotification({ message, user: user.id });
    } catch (error) {
        console.error("Error al guardar la notificación:", error);
        throw error; // Lanzar el error para que pueda ser capturado en la ruta de la API
    }
};

module.exports = { initWebSocket, saveNotification };
