from fastapi.openapi.utils import get_openapi

def custom_openapi(app):
    if app.openapi_schema:
        return app.openapi_schema
    openapi_schema = get_openapi(
        title="Create Tourist Place API",
        version="1.0",
        description="API para crear lugares turísticos",
        routes=app.routes,
    )
    app.openapi_schema = openapi_schema
    return app.openapi_schema
