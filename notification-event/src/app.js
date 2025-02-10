const express = require("express");
const axios = require("axios");
const pool = require("./config/db");
const app = express();
app.use(express.json()); 

// URL del microservicio Validation (asegúrate de que esté corriendo)
const VALIDATION_SERVICE_URL = "http://localhost:8080/buscar";

const validateUserAndPlace = async (username, name) => {
    try {
        const response = await axios.post(VALIDATION_SERVICE_URL, {
            valor_postgres: username,
            valor_mongo: name,
        });

        return response.data.resultado; // Devuelve true si ambos existen
    } catch (error) {
        console.error("Error en la validación:", error.response ? error.response.data : error.message);
        return false;
    }
};

// Ruta para recibir eventos vía WebHook
app.post("/api/webhook", async (req, res) => {
    const { event_name, payload, username, name } = req.body;

    if (!event_name || !payload || !username || !name) {
        return res.status(400).json({ error: "Faltan datos en la solicitud" });
    }

    try {
        // Validar user_id y place_id usando el microservicio Validation
        const isValid = await validateUserAndPlace(username, name);

        if (!isValid) {
            return res.status(404).json({ error: "Usuario o Lugar no encontrado" });
        }

        // Guardar evento en PostgreSQL si la validación fue exitosa
        const query = `
            INSERT INTO events (event_name, payload, user_name, place_name)
            VALUES ($1, $2, $3, $4) RETURNING *;
        `;

        const values = [event_name, JSON.stringify(payload), username, name];

        const result = await pool.query(query, values);
        console.log("✅ Evento almacenado:", result.rows[0]);

        res.status(201).json({ message: "Evento registrado", event: result.rows[0] });
    } catch (error) {
        console.error("Error al procesar el evento:", error);
        res.status(500).json({ error: "Error interno del servidor" });
    }
});

const PORT = process.env.PORT || 5006;
app.listen(PORT, () => {
    console.log(`✅ Servidor WebHook escuchando en el puerto ${PORT}`);
});
