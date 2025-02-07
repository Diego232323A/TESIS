package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Variable global de conexión a la BD
var db *sql.DB

// Carga las variables de entorno
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}
}

// Conecta a la base de datos PostgreSQL
func ConnectDB() {
	LoadEnv()

	var err error
	db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	))

	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo establecer la conexión a la BD:", err)
	}

	fmt.Println("✅ Conectado a PostgreSQL")
}

// Devuelve la conexión a la base de datos
func GetDB() *sql.DB {
	return db
}
