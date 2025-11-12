package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shagabiev/Go-QA-Api/internal/server"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=postgres user=postgres password=postgres dbname=qa_db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect database:", err)
	}

	sqlDB, _ := db.DB()
	goose.SetDialect("postgres")

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		logrus.Fatal("migration failed:", err)
	}

	handler := server.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
