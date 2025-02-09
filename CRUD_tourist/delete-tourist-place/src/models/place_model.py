from pydantic import BaseModel
from typing import Optional

class PlaceModel(BaseModel):
    name: str
    description: str
    location: str
    category: str
    image_url: Optional[str] = None
    