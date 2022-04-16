package main

import (
	"admin-ambassador-checkout/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//user := os.Getenv("MySQL_USER")
	//pass := os.Getenv("MYSQL_PASSWORD")
	//host := os.Getenv("MYSQL_HOST") //Here!!
	//dbname := os.Getenv("MYSQL_DATABASE")
	//connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname) //Fix!!
	//_, err := gorm.Open(mysql.Open(connection))
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,World ")
	})

	app.Listen(":8000")
}
