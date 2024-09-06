// Package circle holds type(struct) and method for circle shape
package circle

// Circle is a public struct but with private field
type Circle struct {
	radius float64
}

// NewCircle is a Constructor to instantiate Circle struct's private fields outside this package
func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

// GetRadius is Getter Method
func (c *Circle) GetRadius() float64 {
	return c.radius
}

// SetRadius is a Setter Method
func (c *Circle) SetRadius(newRadius float64) {
	c.radius = newRadius
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}
