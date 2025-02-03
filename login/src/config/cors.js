const cors = require("cors");

const corsOptions = {
  origin: "*", // Puedes restringirlo a tu frontend
  methods: ["GET", "POST"],
  allowedHeaders: ["Content-Type", "Authorization"]
};

module.exports = cors(corsOptions);
