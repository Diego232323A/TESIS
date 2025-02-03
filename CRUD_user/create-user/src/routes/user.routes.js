const express = require("express");
const { createUser } = require("../controllers/user.controller");

const router = express.Router();

/**
 * @swagger
 * /users:
 *   post:
 *     summary: Crear un nuevo usuario
 *     tags: [Users]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               username:
 *                 type: string
 *                 example: "usuario123"
 *               email:
 *                 type: string
 *                 example: "usuario@email.com"
 *               password:
 *                 type: string
 *                 example: "contraseñaSegura123"
 *     responses:
 *       201:
 *         description: Usuario creado con éxito
 *       400:
 *         description: Datos inválidos
 *       500:
 *         description: Error en el servidor
 */
router.post("/", createUser);

module.exports = router;
