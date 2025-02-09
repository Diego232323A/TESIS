const express = require("express");
const cors = require("cors");
const userRoutes = require("./routes/user.routes");
const swaggerDocs = require("./config/swagger");

require("dotenv").config();

const app = express();
app.use(express.json());
app.use(cors());

app.use("/users", userRoutes);
swaggerDocs(app);
//port isac y karen por siempre
const PORT = process.env.PORT || 4001;
app.listen(PORT, () => {
  console.log(`🚀 Servidor corriendo en http://localhost:${PORT}`);
});