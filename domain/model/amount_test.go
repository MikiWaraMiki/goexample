package domain_model

import "testing"

func TestAmountPriceWhenTragedyAndAudienceLessThan31(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 30,
	}

	amount := NewAmount(&play, &performance, NewUsd())

	expected := 40000

	if result := amount.Price(); result != expected {
		t.Errorf("expected: %v, result: %v", expected, result)
	}
}

func TestAmountPriceWhenTragedyAndAudienceOverThan30(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "tragedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 31,
	}

	amount := NewAmount(&play, &performance, NewUsd())

	expected := 40000 + (1000 * (31 - 30)) // 41000

	if result := amount.Price(); result != expected {
		t.Errorf("expected: %v, result: %v", expected, result)
	}
}

func TestAmountPriceWhenComedyAndAudienceLessThan21(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "comedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 20,
	}

	amount := NewAmount(&play, &performance, NewUsd())

	expected := 30000 + (300 * performance.Audience)

	if result := amount.Price(); result != expected {
		t.Fatalf("expected: %v, result: %v", expected, result)
	}
}

func TestAmountPriceWhenComedyAndAudienceOver20(t *testing.T) {
	play := Play{
		PlayID:   "test",
		Name:     "test",
		TypeName: "comedy",
	}
	performance := Performance{
		PlayID:   "test",
		Audience: 21,
	}

	amount := NewAmount(&play, &performance, NewUsd())

	expected := 30000 + (300 * performance.Audience) + (10000 + 500*(performance.Audience-20))

	if result := amount.Price(); result != expected {
		t.Fatalf("expected: %v, result: %v", expected, result)
	}
}
