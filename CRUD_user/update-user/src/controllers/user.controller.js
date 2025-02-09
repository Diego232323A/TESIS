const pool = require("../db");
const bcrypt = require("bcryptjs");

const updateUser = async (req, res) => {
  const { username } = req.params;
  const { email, password } = req.body;

  try {
    const hashedPassword = await bcrypt.hash(password, 10);
    const result = await pool.query(
      "UPDATE users SET email = $1, password = $2 WHERE username = $3 RETURNING *",
      [email, hashedPassword, username]
    );

    if (result.rowCount === 0) {
      return res.status(404).json({ error: "Usuario no encontrado" });
    }

    res.json({ message: "Usuario actualizado", user: result.rows[0] });
  } catch (error) {
    res.status(500).json({ error: "Error en el servidor" });
  }
};

module.exports = { updateUser };
