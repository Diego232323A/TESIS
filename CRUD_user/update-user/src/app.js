const express = require("express");
const cors = require("./config/cors");
const setupSwagger = require("./config/swagger");
const userRoutes = require("./routes/user.routes");
require("dotenv").config();

const app = express();

app.use(express.json());
app.use(cors);
setupSwagger(app);

app.use("/api", userRoutes);

const PORT = process.env.PORT || 4004;
app.listen(PORT, () => {
  console.log(`✅ Update-User Service corriendo en http://localhost:${PORT}`);
  console.log(`📄 Swagger disponible en http://localhost:${PORT}/api-docs`);
});
