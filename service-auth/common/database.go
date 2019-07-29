package common

import (
	"github.com/jinzhu/gorm"
	"github.com/nekko-ru/api/service-auth/helpers"
	"github.com/nekko-ru/api/service-auth/models"
	log "github.com/sirupsen/logrus"
)

// Database class
type Database struct {
	*gorm.DB
}

// DB contain current connection
var DB *gorm.DB

// Init Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	host := helpers.GetEnv("DB_HOST", "127.0.0.1")
	user := helpers.GetEnv("DB_USER", "postgres")
	dbName := helpers.GetEnv("DB_NAME", "microapi")
	psw := helpers.GetEnv("DB_PSW", "postgres")

	db, err := gorm.Open("postgres", "sslmode=disable host="+host+" user="+user+" dbname="+dbName+" password="+psw)
	if err != nil {
		log.Panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db

	// todo: add auto migration
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Token{})

	log.Info("Database connected")
	return DB
}
