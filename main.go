package main

import (
	"admin-ambassador-checkout/src/database"
	"admin-ambassador-checkout/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // Allow the frontend to get Cookies from the backend
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
