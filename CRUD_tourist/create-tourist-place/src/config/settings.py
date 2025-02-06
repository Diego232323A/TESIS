import os
from dotenv import load_dotenv

load_dotenv()

class Settings:
    PROJECT_NAME: str = "Create Tourist Place API"
    PROJECT_VERSION: str = "1.0"
    
    MONGO_URI: str = os.getenv("MONGO_URI", "mongodb://localhost:27017")
    DB_NAME: str = os.getenv("DB_NAME", "tourism")
    PORT: int = int(os.getenv("PORT", 5001))

settings = Settings()
