const express = require("express");
const client = require("prom-client");
const { saveAlert } = require("./db/db");

const app = express();
app.use(express.json());

// Inicializa métricas
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics();

const requestCount = new client.Counter({
    name: "http_requests_total",
    help: "Número total de solicitudes HTTP",
    labelNames: ["method", "route", "status"],
});

app.use((req, res, next) => {
    res.on("finish", () => {
        requestCount.labels(req.method, req.path, res.statusCode).inc();
    });
    next();
});

// Endpoint de métricas para Prometheus
app.get("/metrics", async (req, res) => {
    res.set("Content-Type", client.register.contentType);
    res.end(await client.register.metrics());
});

// Endpoint para recibir alertas de Prometheus
app.post("/alert", async (req, res) => {
    const alerts = req.body.alerts || [];
    for (let alert of alerts) {
        await saveAlert(alert);
    }
    res.status(200).json({ message: "Alertas guardadas" });
});

const PORT = process.env.PORT || 5009;
app.listen(PORT, () => {
    console.log(`🚀 Microservicio de monitoreo corriendo en el puerto ${PORT}`);
});
