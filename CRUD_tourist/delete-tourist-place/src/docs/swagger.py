from fastapi.openapi.utils import get_openapi

def swagger_config(app):
    def custom_openapi():
        if app.openapi_schema:
            return app.openapi_schema
        openapi_schema = get_openapi(
            title="API - Eliminar Lugar Turístico",
            version="1.0.0",
            description="API para eliminar un lugar turístico",
            routes=app.routes,
        )
        app.openapi_schema = openapi_schema
        return app.openapi_schema
    app.openapi = custom_openapi
