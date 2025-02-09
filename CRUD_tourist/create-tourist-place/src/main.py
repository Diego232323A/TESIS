from fastapi import FastAPI
from fastapi.openapi.docs import get_swagger_ui_html
from src.routes.place_routes import router
from src.docs.swagger import custom_openapi
import uvicorn

app = FastAPI(title="Create Tourist Place API", description="API para crear lugares turísticos", version="1.0")

# Configurar Swagger UI
@app.get("/docs", include_in_schema=False)
async def get_docs():
    return get_swagger_ui_html(openapi_url="/openapi.json", title="API Docs")

app.openapi = lambda: custom_openapi(app)
app.include_router(router, prefix="/api")

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5001)
