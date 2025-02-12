const { Pool } = require("pg");
require("dotenv").config();

const pool = new Pool({
    user: process.env.DB_USER,
    host: process.env.DB_HOST,
    database: process.env.DB_NAME,
    password: process.env.DB_PASSWORD,
    port: process.env.DB_PORT,
  });

const saveAlert = async (alert) => {
    const { labels, startsAt, endsAt, status } = alert;
    const query = `
        INSERT INTO alerts (name, severity, starts_at, ends_at, status) 
        VALUES ($1, $2, $3, $4, $5)
    `;
    await pool.query(query, [
        labels.alertname,
        labels.severity,
        startsAt,
        endsAt,
        status,
    ]);
};

module.exports = { saveAlert };
