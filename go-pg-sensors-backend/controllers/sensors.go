package controllers

import (
	"github.com/claudiumocanu/microservices-having-fun/go-pg-sensors-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Sensor struct {
	gorm.Model
	SensorName          string  `json:"sensorname"`
	SeonsorUUID         string  `json:"seonsoruuid"`
	MeasuringUnit       string  `json:"measuringunit"`
	MinContructiveValue float32 `json:"mincontructivevalue"`
	MaxContructiveValue float32 `json:"maxcontructivevalue"`
	MinNominalValue     float32 `json:"minnominalvalue"`
	MaxNominalValue     float32 `json:"maxnominalvalue"`
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
