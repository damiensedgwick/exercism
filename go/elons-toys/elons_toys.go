package elon

import "fmt"

func (c *Car) Drive() {
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %i meters", c.speed)
}

func (c *Car) DisplayBattery() string {
	battery := c.battery - c.batteryDrain
	return fmt.Sprintf("Battery at %i%", battery)
}

func (c *Car) CanFinish(distance int) bool {}
