const { pool } = require("../config/db");

const createWebhook = async (url) => {
    const query = "INSERT INTO webhooks (url) VALUES ($1) RETURNING *";
    const values = [url];
    const result = await pool.query(query, values);
    return result.rows[0];
};

const getWebhooks = async () => {
    const result = await pool.query("SELECT * FROM webhooks");
    return result.rows;
};

module.exports = { createWebhook, getWebhooks };