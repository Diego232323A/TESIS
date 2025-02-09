const mongoose = require("mongoose");

// Esquema de la notificación
const notificationSchema = new mongoose.Schema({
    message: { type: String, required: true },
    user: { type: Number, required: true },
    timestamp: { type: Date, default: Date.now },
});

const Notification = mongoose.model("Notification", notificationSchema);

module.exports = Notification;
