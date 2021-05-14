package database

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/models"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var env map[string]string
	env, err := godotenv.Read();
	if err != nil {
		log.Fatal("No .env file found")
	}
	envuser     := env["USER"]
	envpassword := env["PASSWORD"]
	envdbname   := env["DBNAME"]

	database, err := gorm.Open(mysql.Open(envuser+":"+envpassword+"@/"+envdbname), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = database

	database.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	
}

