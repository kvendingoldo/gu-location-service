package distance

import (
	"fmt"
	"testing"
)

func TestCalculateDistance_1(t *testing.T) {
	coordinates := []Coordinate{
		{35.12314, 27.64532},
		{39.12355, 27.64532},
		{35.12314, 27.64532},
	}

	got, err := CalculateDistance(coordinates, "km")
	if err != nil {
		t.Errorf("got error: %v", err)

	}

	want := 889.6505930598652
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateDistance_2(t *testing.T) {
	coordinates := []Coordinate{
		{35.12314, 27.64532},
		{39.12355, 27.64538},
	}

	got, err := CalculateDistance(coordinates, "km")
	if err != nil {
		t.Errorf("got error: %v", err)

	}

	want := 444.8252965299326
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateDistance_3(t *testing.T) {
	coordinates := []Coordinate{
		{35.12314, 27.64532},
		{39.12355, 27.64538},
	}

	got, err := CalculateDistance(coordinates, "mi")
	if err != nil {
		t.Errorf("got error: %v", err)

	}

	want := 276.4015393301007
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func ExampleCalculateDistance() {
	coordinates := []Coordinate{
		{35.12314, 27.64532},
		{39.12355, 27.64532},
		{35.12314, 27.64532},
	}
	distance, err := CalculateDistance(coordinates, "km")
	if err != nil {
		// Handle it
	}
	fmt.Println(distance)
	// Output:
	// 889.6505930598652
}
