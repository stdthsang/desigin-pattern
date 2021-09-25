package visitor

import "fmt"

type visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}

type Shape interface {
	GetType() string
	Accept(visitor)
}

type Square struct {
	side int
}

func (s *Square) Accept(v visitor) {
	v.visitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) Accept(v visitor) {
	v.visitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) Accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) GetType() string {
	return "Rectangle"
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Print("Calculating area for square\n")
}

func (a *areaCalculator) visitForCircle(s *Circle) {
	fmt.Print("Calculating area for circle\n")
}

func (a *areaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Print("Calculating area for rectangle\n")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Print("Calculating middle point coordinates for square\n")
}

func (a *middleCoordinates) visitForCircle(c *Circle) {
	fmt.Print("Calculating middle point coordinates for circle\n")
}

func (a *middleCoordinates) visitForrectangle(t *Rectangle) {
	fmt.Print("Calculating middle point coordinates for rectangle\n")
}
