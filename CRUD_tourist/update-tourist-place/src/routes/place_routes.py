from fastapi import APIRouter, HTTPException, status
from src.database.connection import places_collection
from src.models.place_model import PlaceUpdate
from bson import ObjectId

router = APIRouter()

@router.put("/place/{place_id}", response_model=dict)
async def update_place(place_id: str, place_data: PlaceUpdate):
    if not ObjectId.is_valid(place_id):
        raise HTTPException(status_code=400, detail="ID de lugar inválido")

    update_data = {k: v for k, v in place_data.dict().items() if v is not None}

    if not update_data:
        raise HTTPException(status_code=400, detail="No se enviaron datos para actualizar")

    result = places_collection.update_one({"_id": ObjectId(place_id)}, {"$set": update_data})

    print(f"Resultado de la actualización: {result.raw_result}") 

    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Lugar turístico no encontrado en la base de datos")

    return {
        "message": "Lugar turístico actualizado correctamente",
        "matched_count": result.matched_count,
        "modified_count": result.modified_count
    }



