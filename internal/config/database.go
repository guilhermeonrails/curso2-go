package config

import (
	"log"
	"myapi/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Conexão com o Postgres (usando host "db" pois o docker-compose cria essa rede)
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o BD: %v", err)
	}
	DB = db

	err = DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("Erro durante a migração: %v", err)
	}
}
