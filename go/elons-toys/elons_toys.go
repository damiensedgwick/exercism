package elon

func (c *Car) Drive() {
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

func (c *Car) DisplayDistance() string {}

func (c *Car) DisplayBattery() string {}

func (c *Car) CanFinish(distance int) bool {}
