package main

import (
	"testing"
)

func TestAmountForByTragedyAudienceLessThan31(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 30,
	}

	expected := 40000

	if result := AmountFor(&performance, &play); result != expected {
		t.Errorf("AmountFor() = %v, want %v", result, expected)
	}
}

func TestAmoundForByTragedyAudienceOverThan30(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 31,
	}

	expected := 40000 + (1000 * (31 - 30))
	if result := AmountFor(&performance, &play); result != expected {
		t.Errorf("AmountFor() = %v, want %v", result, expected)
	}
}

func TestAmountForByComedyAudienceLessThan21(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "comedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 20,
	}

	expected := 30000 + (300 * 20)
	if result := AmountFor(&performance, &play); result != expected {
		t.Errorf("AmountFor() = %v, want %v", result, expected)
	}
}
func TestAmountForByComedyAudienceOverThan20(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "comedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 21,
	}

	expected := 30000 + 10000 + 500*(21-20) + (300 * 21)
	if result := AmountFor(&performance, &play); result != expected {
		t.Errorf("AmountFor() = %v, want %v", result, expected)
	}
}
