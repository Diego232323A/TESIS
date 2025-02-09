from pydantic import BaseModel
from typing import Optional

class Place(BaseModel):
    name: str
    description: str
    location: str
    category: str
    rating: Optional[float] = None
