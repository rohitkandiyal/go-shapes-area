// Package wheel holds type(struct) and method for wheel
package wheel

import (
	// _ "fmt"

	"github.com/rohitkandiyal/go-shapes-area/packages/circle"
)

// Wheel is a public struct but with private field
type Wheel struct {
	*circle.Circle
	vehicle, material string
}

// NewWheel is a Constructor to instantiate Wheel struct's private fields outside this package
func NewWheel(radius float64, vehicle, material string) *Wheel {
	return &Wheel{circle.NewCircle(radius), vehicle, material}
}

func (w *Wheel) GetName() string {
	return w.vehicle
}

func (w *Wheel) GetCost() float64 {
	cost := 0.0

	if w.material == "aluminium" {
		cost = cost + 20.0
	} else if w.material == "cast-iron" {
		cost = cost + 10.0
	} else {
		cost = cost + 5.0
	}

	if w.Area() < 100.0 {
		cost = cost + 50.0
	} else if w.Area() >= 50.0 {
		cost = cost + 100.0
	}

	return cost

}
