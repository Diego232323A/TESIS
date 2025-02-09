import os
from dotenv import load_dotenv

# Cargar variables del archivo .env
load_dotenv()

class Settings:
    MONGO_URI: str = os.getenv("MONGO_URI")
    DB_NAME: str = os.getenv("DB_NAME")
    COLLECTION_NAME: str = os.getenv("COLLECTION_NAME")
    PORT: int = int(os.getenv("PORT"))

settings = Settings()
