from fastapi import FastAPI
from src.routes.place_routes import router
import uvicorn

app = FastAPI(title="Actualizar Lugar Turístico API", version="1.0")

app.include_router(router, prefix="/api", tags=["Lugares Turísticos"])

if __name__ == "__main__":
    uvicorn.run("src.main:app", host="0.0.0.0", port=5002, reload=True)
