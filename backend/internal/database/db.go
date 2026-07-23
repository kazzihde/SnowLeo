package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/datashelll/SnowLeo/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB ping failed:", err)
	}

	log.Println("Connected to MariaDB successfully")
	return db
}
