package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "host=localhost user=postgres password=12345 dbname=geo-location sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ No se pudo establecer conexión con PostgreSQL:", err)
	}

	fmt.Println("✅ Conectado a PostgreSQL con PostGIS")
}
