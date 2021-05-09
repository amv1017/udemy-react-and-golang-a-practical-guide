package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	var env map[string]string
	env, err := godotenv.Read();
	if err != nil {
		log.Fatal("No .env file found")
	}
	envuser     := env["USER"]
	envpassword := env["PASSWORD"]
	envdbname   := env["DBNAME"]

	db, err := gorm.Open(mysql.Open(envuser+":"+envpassword+"@/"+envdbname), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	log.Println(db)
}

