const express = require("express");
const { login } = require("../controllers/auth.controller");
const validateLogin = require("../middleware/validateLogin");

const router = express.Router();

router.post("/login", validateLogin, login);

module.exports = router;
