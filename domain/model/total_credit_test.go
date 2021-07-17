package domain_model

import (
	"testing"
)

func TestAddCredit(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 30,
	}

	totalCredit := NewTotalCredit()

	credit := NewCredit(&performance, &play)

	totalCredit.AddCredit(credit)

	if result := totalCredit.Length(); result != 1 {
		t.Errorf("expected 1, result %v", result)
	}
}

func TestVolume(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 31,
	}

	totalCredit := NewTotalCredit()

	credit := NewCredit(&performance, &play)

	totalCredit.AddCredit(credit)
	totalCredit.AddCredit(credit)

	expected := credit.Volume() * 2

	if result := totalCredit.Volume(); result != expected {
		t.Errorf("expected: %v, result: %v", expected, result)
	}
}
