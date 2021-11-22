package lasagna

import "testing"

type lasagnaTests struct {
	name     string
	layers   int
	time     int
	expected int
}

func TestOvenTime(t *testing.T) {

	tt := lasagnaTests{
		name:     "Calculates how many minutes the lasagna should be in the oven",
		layers:   0,
		time:     40,
		expected: 40,
	}

	if got := OvenTime; got != tt.expected {
		t.Errorf("OvenTime(%d) = %d; want %d", tt.expected, got, tt.expected)
	}

}

func TestRemainingOvenTime(t *testing.T) {

	tt := lasagnaTests{
		name:     "Calculates the remaining oven time",
		time:     30,
		expected: 10,
	}

	if got := RemainingOvenTime(tt.time); got != tt.expected {
		t.Errorf("RemainingOvenTime(%d) = %d; want %d", tt.time, got, tt.expected)
	}
}

func TestPreparationTime(t *testing.T) {

	tt := lasagnaTests{
		name:     "Calculates the preparation time depending on the amount of layers",
		time:     2, // time = two minutes per later
		layers:   5,
		expected: 10,
	}

	if got := PreparationTime(tt.layers); got != tt.expected {
		t.Errorf("PreparationTime(%d) = %d; want %d", tt.layers, got, tt.expected)

	}

}

func TestElapseTime(t *testing.T) {

	tt := lasagnaTests{
		name:     "Calculates the elapse time",
		time:     20, // time = time in the oven
		layers:   2,
		expected: 24,
	}

	if got := ElapsedTime(tt.layers, tt.time); got != tt.expected {
		t.Errorf("ElapsedTIme(%d, %d) = %d; want %d", tt.layers, tt.time, got, tt.expected)
	}

}
