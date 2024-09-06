// Package chassis holds type(struct) and method for chassis
package chassis

import (
	// _ "fmt"

	"github.com/rohitkandiyal/go-shapes-area/packages/rectangle"
)

// Chassis is a public struct but with private fields
type Chassis struct {
	rectangle.Rrectangle
	vehicle, material string
}

// NewChassis is a Constructor to instantiate Wheel struct's private fields outside this package
func NewChassis(length, breadth float64, vehicle, material string) Chassis {
	return Chassis{rectangle.NewRectangle(length, breadth), vehicle, material}
}

func (c Chassis) GetName() string {
	return c.vehicle
}

func (c Chassis) GetCost() float64 {
	cost := 0.0

	if c.material == "aluminium" {
		cost = cost + 10020.0
	} else if c.material == "cast-iron" {
		cost = cost + 10010.0
	} else {
		cost = cost + 10005.0
	}

	if c.Area() < 1000.0 {
		cost = cost + 10050.0
	} else if c.Area() >= 1000.0 {
		cost = cost + 10100.0
	}

	return cost

}
