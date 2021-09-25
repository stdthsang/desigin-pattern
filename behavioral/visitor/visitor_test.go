package visitor

func ExampleAreaCalculator() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
	// Output:
	// Calculating area for square
	// Calculating area for circle
	// Calculating area for rectangle
}

func ExampleMiddleCoordinates() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
	// Output:
	// Calculating middle point coordinates for square
	// Calculating middle point coordinates for circle
	// Calculating middle point coordinates for rectangle
}
