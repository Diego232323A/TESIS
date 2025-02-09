const express = require("express");
const { deleteUser } = require("../controllers/user.controller");
const authenticateToken = require("../middleware/authMiddleware");

const router = express.Router();

/**
 * @swagger
 * /api/user/{username}:
 *   delete:
 *     summary: Eliminar un usuario por username
 *     tags: [Usuarios]
 *     parameters:
 *       - in: path
 *         name: username
 *         required: true
 *         description: Nombre del usuario a eliminar
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Usuario eliminado exitosamente
 *       404:
 *         description: Usuario no encontrado
 *       401:
 *         description: No autorizado
 */
router.delete("/user/:username", authenticateToken, deleteUser);

module.exports = router;
