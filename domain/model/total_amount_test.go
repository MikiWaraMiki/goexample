package domain_model

import "testing"

func TestAddAmount(t *testing.T) {
	totalAmount := NewTotalAmount()

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

	totalAmount.AddAmount(amount)

	if length := len(totalAmount.AllAmount()); length != 1 {
		t.Errorf("expected 1 but was %v", length)
	}
}

func TestPrice(t *testing.T) {
	totalAmount := NewTotalAmount()

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

	totalAmount.AddAmount(amount)
	totalAmount.AddAmount(amount)

	expected := amount.Price() * 2

	if result := totalAmount.Price(); result != expected {
		t.Errorf("expected %v, result %v", expected, result)
	}
}
