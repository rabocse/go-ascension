package Annalyns_Infiltration

import "testing"

// Type to contain the input and outputs needed for the tests.
type annalynsTests struct {
	name             string
	knightIsAwake    bool
	archerIsAwake    bool
	prisonerIsAwake  bool
	dogIsWithAnnalyn bool
	expected         bool
}

// 1. Check if a fast attack can be made test.
func TestCanFastAttack(t *testing.T) {

	tt := annalynsTests{
		name:          "Decide if a fast attack can be made",
		knightIsAwake: true,
		expected:      true,
	}

	if got := CanFastAttack(tt.knightIsAwake); got != tt.expected {
		t.Errorf("CanFastAttack(%v) = %v; but the expected is %v", tt.knightIsAwake, got, tt.expected)
	}

}

// 2. Check if the group can be spied upon test.
func TestCanSpy(t *testing.T) {

	tt := annalynsTests{
		name:            "Decide if spying on is the best option to make",
		knightIsAwake:   true,
		archerIsAwake:   false,
		prisonerIsAwake: false,
		expected:        true,
	}

	if got := CanSpy(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake); got != tt.expected {
		t.Errorf("CanSpy(%v, %v, %v) = %v; but the expected is %v", tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, got, tt.expected)
	}

}

// 3. Check if the prisoner can be signaled test.
func TestCanSignalPrisoner(t *testing.T) {

	tt := annalynsTests{
		name:            "Decide if prisoner can be signaled",
		knightIsAwake:   true,
		archerIsAwake:   false,
		prisonerIsAwake: true,
		expected:        true,
	}

	if got := CanSignalPrisoner(tt.prisonerIsAwake, tt.archerIsAwake); got != tt.expected {
		t.Errorf("CanSignalPrisoner(%v, %v) = %v; but the expected is %v", tt.prisonerIsAwake, tt.archerIsAwake, got, tt.expected)
	}

}

// 4. Check if the prisoner can be freed test.
func TestCanFreePrisoner(t *testing.T) {

	tt := annalynsTests{
		name:             "Decide if the prisoner can be freed",
		knightIsAwake:    true,
		archerIsAwake:    false,
		prisonerIsAwake:  true,
		dogIsWithAnnalyn: true,
		expected:         true,
	}

	if got := CanFreePrisoner(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.dogIsWithAnnalyn); got != tt.expected {
		t.Errorf("CanFreePrisoner(%v, %v) = %v; but the expected is %v", tt.archerIsAwake, tt.dogIsWithAnnalyn, got, tt.expected)
	}
}
