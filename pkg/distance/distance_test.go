package distance

import (
	"fmt"
	"testing"
)

func TestCalculateDistanceInKm_1(t *testing.T) {
	coordinates := []string{
		"35.12314, 27.64532",
		"39.12355, 27.64538",
		"35.12314, 27.64532",
	}

	got, err := CalculateDistanceInKm(coordinates)
	if err != nil {
		t.Errorf("got error: %v", err)

	}

	want := 889.6505930598652
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateDistanceInKm_2(t *testing.T) {
	coordinates := []string{
		"35.12314, 27.64532",
		"39.12355, 27.64538",
	}

	got, err := CalculateDistanceInKm(coordinates)
	if err != nil {
		t.Errorf("got error: %v", err)

	}

	want := 444.8252965299326
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateDistanceInMiles_1(t *testing.T) {
	coordinates := []string{
		"35.12314, 27.64532",
		"39.12355, 27.64538",
	}

	got, err := CalculateDistanceInMiles(coordinates)
	if err != nil {
		t.Errorf("got error: %v", err)

	}

	want := 276.4015393301007
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func ExampleCalculateDistanceInKm() {
	coordinates := []string{
		"35.12314, 27.64532",
		"39.12355, 27.64538",
		"35.12314, 27.64532",
	}
	distance, err := CalculateDistanceInKm(coordinates)
	if err != nil {
		// Handle it
	}
	fmt.Println(distance)
	// Output:
	// 889.6505930598652
}

func ExampleCalculateDistanceInMiles() {
	coordinates := []string{
		"35.12314, 27.64532",
		"39.12355, 27.64538",
	}
	distance, err := CalculateDistanceInMiles(coordinates)
	if err != nil {
		// Handle it
	}
	fmt.Println(distance)
	// Output:
	// 276.4015393301007
}
