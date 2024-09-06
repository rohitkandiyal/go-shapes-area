// Package main
package main

import (
	"fmt"
	// "github.com/rohitkandiyal/go-shapes-area/pkg/circle"
	"github.com/rohitkandiyal/go-shapes-area/packages/chassis"
	"github.com/rohitkandiyal/go-shapes-area/packages/rectangle"
	"github.com/rohitkandiyal/go-shapes-area/packages/wheel"
	_ "github.com/spf13/cobra"
)

type carPart interface {
	GetName() string
	GetCost() float64
}

func main() {
	// Declaring public struct directly: But use constructor even if all fields of a struct are public
	myrect := rectangle.Rectangle{
		Length:  6.50,
		Breadth: 8.50,
	}
	fmt.Println("Length:", myrect.Length, "Breadth:", myrect.Breadth)

	// Accessing struct with private values via Constructor/Setters/Getters
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Accessing struct with private fields via Constructors/Setters/Getters")
	myrect1 := rectangle.NewRectangle(15.50, 22.50)
	fmt.Println("Length:", myrect1.Length, "Breadth:", myrect1.GetBreadth())
	myrect1.SetBreadth(19.50)
	fmt.Println("After change via SETTER")
	fmt.Println("Length:", myrect1.Length, "Breadth:", myrect1.GetBreadth())
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")

	// COMPOSITION: embedded struct -- here we need to calculate the price of wheel including tax(calculated as per it's diameter)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("COMPOSITION- Calculating Car Cost = Wheel cost + Chassis cost")

	// No INTERFACE USAGE
	carParts := []*wheel.Wheel{
		wheel.NewWheel(20.80, "jeep", "cast-iron"),
		wheel.NewWheel(15.20, "kia", "aluminium"),
		wheel.NewWheel(25.60, "defender", "cast-iron"),
	}

	for _, b := range carParts {
		fmt.Println("Radius of ", b.GetName(), "wheel's is", b.GetRadius(), "and the cost is", b.GetCost())
	}

	// USING INTERFACE

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("COMPOSITION with Interface- Calculating Car Cost = Wheel cost + Chassis cost")

	carParts1 := []carPart{
		wheel.NewWheel(20.80, "jeep", "cast-iron"),
		chassis.NewChassis(200.80, 150.80, "jeep", "cast-iron"),
		wheel.NewWheel(15.20, "kia", "aluminium"),
		chassis.NewChassis(200.80, 150.80, "kia", "aluminium"),
		wheel.NewWheel(25.60, "defender", "cast-iron"),
		chassis.NewChassis(200.80, 150.80, "defender", "cast-iron"),
	}

	for _, b := range carParts1 {
		// Type assertion needed here
		switch item := b.(type) {
		case *wheel.Wheel:
			fmt.Println("Cost of WHEEL of ", item.GetName(), "is", item.GetCost())
		case chassis.Chassis:
			fmt.Println("Cost of CHASSIS of ", item.GetName(), "is", item.GetCost())
		default:
			fmt.Println("this is default in switch case")
		}
	}

}
