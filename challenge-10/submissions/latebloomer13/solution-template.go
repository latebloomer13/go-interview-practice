// Package challenge10 contains the solution for Challenge 10.
package challenge10

import (
	"fmt"
	"errors"
	"math"
	"os"
	// Add any necessary imports here
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer // Includes String() string method
}

// Rectangle represents a four-sided shape with perpendicular sides
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle creates a new Rectangle with validation
func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
	    return nil, errors.New("not positive sides of rectangle")
	}
	return &Rectangle{Width: width, Height: height}, nil
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	// TODO: Implement area calculation
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return (r.Width+r.Height)*2
}

// String returns a string representation of the rectangle
func (r *Rectangle) String() string {

	return fmt.Sprintf("Rectangle height=%v  width=%v", r.Height, r.Width)
}

// Circle represents a perfectly round shape
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle with validation
func NewCircle(radius float64) (*Circle, error) {
	// TODO: Implement validation and construction
	if radius <= 0 {
	    return nil, errors.New("not positive radius")
	}
	return &Circle{Radius: radius}, nil
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	// TODO: Implement area calculation
	return math.Pi*c.Radius*c.Radius
}

// Perimeter calculates the circumference of the circle
func (c *Circle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return math.Pi*c.Radius*2
}

// String returns a string representation of the circle
func (c *Circle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf(" Circle  radius=%v", c.Radius)
}

// Triangle represents a three-sided polygon
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// NewTriangle creates a new Triangle with validation
func NewTriangle(a, b, c float64) (*Triangle, error) {
	// TODO: Implement validation and construction
	if a <= 0 || b <= 0 || c <= 0 {
	    return nil, errors.New("side of triangle is not positive")
	}
	if a >= b+c || b >= a+c || c >= a+b {
	    return nil, errors.New("triangle rule is not true")
	}
	return &Triangle{SideA: a, SideB: b, SideC: c}, nil
}

// Area calculates the area of the triangle using Heron's formula
func (t *Triangle) Area() float64 {
	// TODO: Implement area calculation using Heron's formula
	semip := t.Perimeter() / 2
	
	return math.Sqrt(semip*(semip-t.SideA)*(semip-t.SideB)*(semip-t.SideC))
}

// Perimeter calculates the perimeter of the triangle
func (t *Triangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return t.SideA + t.SideB + t.SideC
}

// String returns a string representation of the triangle
func (t *Triangle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Triangle sides=%v %v %v", t.SideA, t.SideB, t.SideC)
}

// ShapeCalculator provides utility functions for shapes
type ShapeCalculator struct{}

// NewShapeCalculator creates a new ShapeCalculator
func NewShapeCalculator() *ShapeCalculator {
	// TODO: Implement constructor
	return nil
}

// PrintProperties prints the properties of a shape
func (sc *ShapeCalculator) PrintProperties(s Shape) {
	// TODO: Implement printing shape properties
	fmt.Println(s)
}

// TotalArea calculates the sum of areas of all shapes
func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {
	// TODO: Implement total area calculation
	var result float64
	for _, shape := range shapes {
	    result += shape.Area()
	}
	return result
}

// LargestShape finds the shape with the largest area
func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	// TODO: Implement finding largest shape
	if len(shapes) == 0 {
        os.Exit(1)
	}
	var result Shape = shapes[0]
	var max_area float64 = shapes[0].Area()
	
	for i := 1; i < len(shapes); i++ {
	    tmp_area := shapes[i].Area()
	    if tmp_area > max_area {
	        max_area = tmp_area
	        result = shapes[i]
	    }
	}
	return result
}

// SortByArea sorts shapes by area in ascending or descending order
func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	// TODO: Implement sorting shapes by area
// 	result := make([]Shape, len(shapes), len(shapes))
    fmt.Println(shapes)
	for i := 0; i < len(shapes); i++ {
        for j := 0; j < len(shapes)-1-i; j++ {
            if ascending {
                if shapes[j].Area() > shapes[j+1].Area() {
                    shapes[j], shapes[j+1] = shapes[j+1], shapes[j]
                }
            } else {
                if shapes[j].Area() < shapes[j+1].Area() {
                    shapes[j], shapes[j+1] = shapes[j+1], shapes[j]
                }
            }
        }
	} 
    fmt.Println(shapes)
	return shapes
} 