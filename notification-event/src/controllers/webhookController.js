const pool = require("../config/db");

// Función para recibir WebHook y almacenarlo en PostgreSQL
const handleWebhook = async (req, res) => {
  try {
    const { event_name, payload, user_id, place_id } = req.body;

    // Insertar el evento en la base de datos
    const query = `
      INSERT INTO events (event_name, payload, user_id, place_id)
      VALUES ($1, $2, $3, $4) RETURNING *;
    `;
    const values = [event_name, JSON.stringify(payload), user_id, place_id];

    const result = await pool.query(query, values);
    console.log("✅ Evento almacenado:", result.rows[0]);

    res.status(201).json({ message: "Evento recibido y almacenado", event: result.rows[0] });
  } catch (error) {
    console.error("❌ Error al procesar WebHook:", error);
    res.status(500).json({ error: "Error interno del servidor" });
  }
};

module.exports = { handleWebhook };
