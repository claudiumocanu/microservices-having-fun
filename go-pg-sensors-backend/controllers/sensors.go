package controllers

import (
	"fmt"

	"github.com/claudiumocanu/microservices-having-fun/go-pg-sensors-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Sensor struct {
	gorm.Model
	SensorName          string  `json:"sensorName"`
	SeonsorUUID         string  `json:"sensorUUID"`
	MeasuringUnit       string  `json:"measuringUnit"`
	MinContructiveValue float32 `json:"minContructiveValue"`
	MaxContructiveValue float32 `json:"maxContructiveValue"`
	MinNominalValue     float32 `json:"minNominalValue"`
	MaxNominalValue     float32 `json:"maxNominalValue"`
	Geolocation         string  `json:"geolocation"`
}

func GetSensors(c *fiber.Ctx) error {
	db := database.DBConn
	var sensors []Sensor
	db.Find(&sensors)
	c.JSON(sensors)
	return nil
}
func GetSensor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var sensor Sensor
	db.Find(&sensor, id)
	c.JSON(sensor)
	return nil
}
func PostSensor(c *fiber.Ctx) error {
	db := database.DBConn
	sensor := new(Sensor)
	if err := c.BodyParser(sensor); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	fmt.Println(sensor)
	if sensor.SensorName == "" || sensor.SeonsorUUID == "" {
		return c.Status(422).SendString("Sensor name and UUID are mandatory fields")
	}
	db.Create(&sensor)
	c.JSON(sensor)
	return nil
}
func DeleteSensor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var sensor Sensor
	db.First(&sensor, id)
	if sensor.ID == 0 {
		return c.Status(500).SendString("Sensor not found")
	}
	db.Delete(&sensor)
	c.SendString("Sensor successfully deleted")
	return nil
}
