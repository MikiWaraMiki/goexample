package domain_model

import (
	"math"
	"testing"
)

func TestVolumeWhenLTragedyAndAudienceLessThan31(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 30,
	}

	credit := NewCredit(&performance, &play)

	expected := 0 // int(math.Max(float64(performance.Audience-30), 0))

	if result := credit.Volume(); result != expected {
		t.Fatalf("expected: %v, result: %v", expected, result)
	}
}

func TestVolumeWhenLTragedyAndAudienceOver30(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 31,
	}

	credit := NewCredit(&performance, &play)

	expected := 1

	if result := credit.Volume(); result != expected {
		t.Fatalf("expected: %v, result: %v", expected, result)
	}
}

func TestVolumeWhenComedy(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "comedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 30,
	}

	credit := NewCredit(&performance, &play)

	expected := 0 + int(math.Trunc(float64(30)/5))

	if result := credit.Volume(); result != expected {
		t.Fatalf("expected: %v, result: %v", expected, result)
	}
}
