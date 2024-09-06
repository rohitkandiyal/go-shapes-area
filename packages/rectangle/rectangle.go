// Package rectangle is holds type(struct) and method for rectangle shape
package rectangle

// Rectangle ... Public
type Rectangle struct {
	Length, Breadth float64
}

// Rrectangle is a public struct but with private and public field
type Rrectangle struct {
	Length, breadth float64
}

// NewRectangle is a Constructor to instantiate Rrectangle struct's private fields outside this package
func NewRectangle(length, breadth float64) Rrectangle {
	return Rrectangle{length, breadth}
}

// GetBreadth is Getter Method
func (r Rrectangle) GetBreadth() float64 {
	return r.breadth
}

// SetBreadth is a Setter Method
func (r Rrectangle) SetBreadth(newBreadth float64) {
	r.breadth = newBreadth
}

func (r Rrectangle) Area() float64 {
	return r.Length * r.breadth
}

func (r Rrectangle) Perimeter() float64 {
	return 2 * (r.breadth + r.breadth)
}
