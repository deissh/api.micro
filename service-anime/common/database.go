package common

import (
	"github.com/jinzhu/gorm"
	"github.com/nekko-ru/api/service-anime/models"
	"github.com/nekko-ru/api/service-auth/helpers"
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
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db

	return DB
}

// Migrate all needed tables
func Migrate() {
	DB.AutoMigrate(&models.Anime{})
	DB.AutoMigrate(&models.Translator{})
}

// Using this function to get a connection, you can create your connection pool here.
//func GetDB() *gorm.DB {
//	return DB
//}
