from fastapi import APIRouter, Query
from src.config.settings import collection
from bson import ObjectId

router = APIRouter(prefix="/api", tags=["Places"])

@router.get("/places")
async def get_places(
    name: str = Query(None, description="Buscar por nombre"),
    location: str = Query(None, description="Buscar por ubicación"),
    category: str = Query(None, description="Filtrar por categoría")
):
    query = {}
    
    if name:
        query["name"] = {"$regex": name, "$options": "i"}  # Búsqueda insensible a mayúsculas
    if location:
        query["location"] = {"$regex": location, "$options": "i"}
    if category:
        query["category"] = category

    places = list(collection.find(query, {"_id": 1, "name": 1, "location": 1, "category": 1}))

    if not places:
        return {"message": "No se encontraron lugares turísticos"}

    # Convertir ObjectId a string
    for place in places:
        place["_id"] = str(place["_id"])

    return places
