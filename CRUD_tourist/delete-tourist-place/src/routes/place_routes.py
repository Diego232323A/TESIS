from fastapi import APIRouter, HTTPException
from bson import ObjectId
from src.database.connection import places_collection

router = APIRouter()

@router.delete("/place/{place_id}")
async def delete_place(place_id: str):
    if not ObjectId.is_valid(place_id):
        raise HTTPException(status_code=400, detail="ID inválido")

    place_object_id = ObjectId(place_id)

    result = places_collection.delete_one({"_id": place_object_id})

    if result.deleted_count == 0:
        raise HTTPException(status_code=404, detail="Lugar turístico no encontrado o no eliminado")

    return {"message": "Lugar turístico eliminado correctamente"}
