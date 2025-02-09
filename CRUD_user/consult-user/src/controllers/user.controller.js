const pool = require("../db");

const getUserByUsername = async (req, res) => {
  const { username } = req.params;

  try {
    // Buscar usuario en la base de datos
    const result = await pool.query("SELECT id, username, email, created_at FROM users WHERE username = $1", [username]);

    if (result.rows.length === 0) {
      return res.status(404).json({ message: "Usuario no encontrado" });
    }

    res.json(result.rows[0]);
  } catch (error) {
    res.status(500).json({ message: "Error interno del servidor" });
  }
};

module.exports = { getUserByUsername };
