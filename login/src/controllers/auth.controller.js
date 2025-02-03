const pool = require("../db");
const bcrypt = require("bcryptjs");
const { generateToken } = require("../config/jwt");

const login = async (req, res) => {
  const { email, password } = req.body;

  try {
    // Buscar usuario en la base de datos
    const result = await pool.query("SELECT * FROM users WHERE email = $1", [email]);
    if (result.rows.length === 0) {
      return res.status(401).json({ message: "Credenciales incorrectas" });
    }

    const user = result.rows[0];

    // Verificar contraseña
    const isValidPassword = await bcrypt.compare(password, user.password);
    if (!isValidPassword) {
      return res.status(401).json({ message: "Credenciales incorrectas" });
    }

    // Generar token JWT
    const token = generateToken(user);
    res.json({ message: "Inicio de sesión exitoso", token });
  } catch (error) {
    res.status(500).json({ message: "Error interno del servidor" });
  }
};

module.exports = { login };
