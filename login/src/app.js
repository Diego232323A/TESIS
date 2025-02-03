const express = require("express");
const cors = require("./config/cors");
const authRoutes = require("./routes/auth.routes");
const swaggerUi = require("swagger-ui-express");
const swaggerSpec = require("./config/swagger");

require("dotenv").config();

const app = express();
app.use(express.json());
app.use(cors);
app.use("/auth", authRoutes);
app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerSpec));

const PORT = process.env.PORT || 4002;
app.listen(PORT, () => {
  console.log(`✅ Auth Service corriendo en http://localhost:${PORT}`);
});
