const express = require("express");
const { getUserByUsername } = require("../controllers/user.controller");
const authenticateToken = require("../middleware/authMiddleware");

const router = express.Router();

/**
 * @swagger
 * /api/user/{username}:
 *   get:
 *     summary: Obtener usuario por username
 *     tags: [Usuarios]
 *     parameters:
 *       - in: path
 *         name: username
 *         required: true
 *         description: Nombre de usuario a buscar
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Usuario encontrado
 *       404:
 *         description: Usuario no encontrado
 *       401:
 *         description: No autorizado
 */
router.get("/user/:username", authenticateToken, getUserByUsername);

module.exports = router;
