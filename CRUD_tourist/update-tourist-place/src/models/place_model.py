from pydantic import BaseModel, Field
from typing import Optional

class PlaceUpdate(BaseModel):
    name: Optional[str] = Field(None, example="Nuevo Nombre del Lugar")
    description: Optional[str] = Field(None, example="Descripción actualizada del lugar")
    location: Optional[str] = Field(None, example="Quito, Ecuador")
    category: Optional[str] = Field(None, example="Montaña")
