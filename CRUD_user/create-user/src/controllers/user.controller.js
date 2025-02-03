const pool = require("../db");
const Joi = require("joi");
const bcrypt = require("bcryptjs");

/**
 * Validación de datos con Joi
 */
const userSchema = Joi.object({
  username: Joi.string().min(3).max(50).required(),
  email: Joi.string().email().required(),
  password: Joi.string().min(6).required(),
});

/**
 * Crear usuario en la base de datos
 */
const createUser = async (req, res) => {
  const { username, email, password } = req.body;
  const hashedPassword = await bcrypt.hash(password, 10);

  // Validar datos
  const { error } = userSchema.validate({ username, email, password });
  if (error) return res.status(400).json({ message: error.details[0].message });

  try {
    const result = await pool.query(
      "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
      [username, email, hashedPassword]
    );
    res.status(201).json({ message: "Usuario creado", userId: result.rows[0].id });
  } catch (err) {
    res.status(500).json({ message: "Error al crear usuario", error: err.message });
  }
};

module.exports = { createUser };
