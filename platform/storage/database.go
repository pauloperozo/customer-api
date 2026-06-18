package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb(models ...interface{}) *gorm.DB {

	var err error

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var dialecto = sqlite.Open("file::memory:?cache=shared")
	DB, err = gorm.Open(dialecto, config)
	if err != nil {
		log.Fatalf("Error en conectar a la db: %v", err)
	}

	log.Println("Conexion OK ")

	err = DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Error en migrar la db: %v", err)
	}

	return DB
}
