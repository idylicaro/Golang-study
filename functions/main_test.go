package functions

import "testing"

func TestPerimeterFunction(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {

	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if result != expected {
			t.Errorf("got %g want %g", result, expected)
		}
	}
	t.Run("rectangle of width 12 and height 6", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		verifyArea(t, rectangle, 72.0)

	})

	t.Run("circle of radius 10", func(t *testing.T) {
		circle := Circle{10.0}
		verifyArea(t, circle, 314.1592653589793)
	})
}

// Table Driven Testing
func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, expected: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expected {
				t.Errorf("got %g want %g", got, tt.expected)
			}
		})
	}
}
