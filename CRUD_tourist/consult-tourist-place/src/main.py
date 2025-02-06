from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from src.routes.place_routes import router
import uvicorn

app = FastAPI(
    title="Microservicio de Consulta de Lugares Turísticos",
    description="Permite buscar lugares turísticos con filtros avanzados.",
    version="1.0.0"
)

# Habilitar CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Incluir rutas
app.include_router(router)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5004, reload=True)
