from fastapi import FastAPI
from src.routes.place_routes import router
from src.docs.swagger import swagger_config

app = FastAPI()

# Configurar documentación Swagger
swagger_config(app)

# Registrar las rutas
app.include_router(router, prefix="/api", tags=["Lugares Turísticos"])

# Mensaje de bienvenida
@app.get("/", response_model=dict)
async def root():
    return {"message": "API para eliminar lugares turísticos"}
