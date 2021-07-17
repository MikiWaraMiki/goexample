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

// Find
func TestFindWhenIncludePlay(t *testing.T) {
	plays := []Play{
		{
			PlayID:   "test",
			Name:     "test",
			TypeName: "test",
		},
	}
	result, _ := Find("test", plays)

	if result == nil {
		t.Errorf("Find() = nil, want play object")
	}
}
func TestFindWhenNotIncludePlay(t *testing.T) {
	plays := []Play{
		{
			PlayID:   "test",
			Name:     "test",
			TypeName: "test",
		},
	}
	result, err := Find("notfound", plays)

	if err == nil {
		t.Errorf("Find()= %v, want nil", result)
	}
}
