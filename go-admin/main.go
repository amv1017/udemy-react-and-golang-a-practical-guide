package main

import (
    "fmt"
    "log"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/joho/godotenv"
)

func main() {

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
    fmt.Println(db)

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    log.Fatal(app.Listen(":8080"))

}
