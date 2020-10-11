package main

import (
	"github.com/claudiumocanu/microservices-having-fun/go-pg-sensors-backend/controllers"
	"github.com/claudiumocanu/microservices-having-fun/go-pg-sensors-backend/database"

	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func homepage(c *fiber.Ctx) error {
	c.SendString("Follow the white rabbit.")
	return nil
}

func setupRoutes(app *fiber.App) {
	app.Get("/", homepage)
	app.Get("/api/v1/sensor", controllers.GetSensors)
	app.Get("/api/v1/sensor/:id", controllers.GetSensor)
	app.Post("/api/v1/sensor/", controllers.PostSensor)
	app.Delete("/api/v1/sensor/:id", controllers.DeleteSensor)
}

func initDbConn() {
	var err error
	var connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))

	database.DBConn, err = gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	fmt.Println("Database connection successfully opened")
	// Creates sensor table in postgres
	database.DBConn.AutoMigrate(&controllers.Sensor{})
	fmt.Println("Sensors table initialized")

}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "",
	}))

	initDbConn()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
