from fastapi import APIRouter, HTTPException
from src.config.db import places_collection
from src.models.place import Place
from src.schemas.place_schema import place_serializer
from bson import ObjectId

router = APIRouter()

@router.post("/place", response_model=dict, tags=["Lugares Turísticos"])
async def create_place(place: Place):
    """
    Crea un nuevo lugar turístico
    """
    new_place = places_collection.insert_one(place.dict())
    created_place = places_collection.find_one({"_id": new_place.inserted_id})
    return {"message": "Lugar turístico creado", "place": place_serializer(created_place)}
