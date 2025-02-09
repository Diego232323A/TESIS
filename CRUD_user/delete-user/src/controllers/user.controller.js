const pool = require("../db");

const deleteUser = async (req, res) => {
  const { username } = req.params;

  try {
    const result = await pool.query(
      "DELETE FROM users WHERE username = $1 RETURNING *",
      [username]
    );

    if (result.rowCount === 0) {
      return res.status(404).json({ error: "Usuario no encontrado" });
    }

    res.json({ message: "Usuario eliminado exitosamente" });
  } catch (error) {
    res.status(500).json({ error: "Error en el servidor" });
  }
};

module.exports = { deleteUser };

