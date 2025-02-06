def place_serializer(place) -> dict:
    return {
        "id": str(place["_id"]),
        "name": place["name"],
        "description": place["description"],
        "location": place["location"],
        "category": place["category"],
        "rating": place.get("rating", None),
    }
