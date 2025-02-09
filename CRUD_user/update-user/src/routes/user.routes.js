const express = require("express");
const { updateUser } = require("../controllers/user.controller");
const authenticateToken = require("../middleware/authMiddleware");

const router = express.Router();

/**
 * @swagger
 * /api/user/{username}:
 *   put:
 *     summary: Actualizar un usuario por username
 *     tags: [Usuarios]
 *     parameters:
 *       - in: path
 *         name: username
 *         required: true
 *         description: Nombre del usuario a actualizar
 *         schema:
 *           type: string
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               email:
 *                 type: string
 *                 example: "nuevoemail@email.com"
 *               password:
 *                 type: string
 *                 example: "nuevacontraseña"
 *     responses:
 *       200:
 *         description: Usuario actualizado
 *       404:
 *         description: Usuario no encontrado
 *       401:
 *         description: No autorizado
 */
router.put("/user/:username", authenticateToken, updateUser);

module.exports = router;
