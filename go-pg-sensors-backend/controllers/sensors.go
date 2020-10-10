package controllers

import (
	"fmt"

	"github.com/claudiumocanu/microservices-having-fun/go-pg-sensors-backend/database"
	"github.com/gofiber/fiber"
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

func GetSensors(c *fiber.Ctx) {
	db := database.DBConn
	var sensors []Sensor
	db.Find(&sensors)
	c.JSON(sensors)
}
func GetSensor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var sensor Sensor
	db.Find(&sensor, id)
	c.JSON(sensor)
}
func PostSensor(c *fiber.Ctx) {
	db := database.DBConn
	sensor := new(Sensor)
	if err := c.BodyParser(sensor); err != nil {
		c.Status(503).Send(err)
		return
	}
	if sensor.SensorName == "" || sensor.SeonsorUUID == "" {
		c.Status(422).Send("Name and UUID are mandatory fileds for a sensor")
		return
	}
	db.Create(&sensor)
	c.JSON(sensor)
}
func DeleteSensor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var sensor Sensor
	db.First(&sensor, id)
	if sensor.ID == 0 {
		c.Status(500).Send(fmt.Sprintf("Sensor not found with ID=%s", id))
		return
	}
	db.Delete(&sensor)
	c.Send("Sensor successfully deleted")
}
