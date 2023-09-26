package database

import (
	"fmt"
	"strconv"

	"github.com/agustfricke/fiber-jwt-auth/config"
	"github.com/agustfricke/fiber-jwt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("Error al analizar el puerto de la base de datos")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error al conectar la base de datos")
	}
	fmt.Println("Conexión abierta a la base de datos")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Base de datos migrada")
}
