package domain_service

import (
	"testing"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

func TestCalcTotalSuccess(t *testing.T) {
	plays := []Play{
		{
			PlayID:   "test",
			Name:     "test",
			TypeName: "tragedy",
		},
		{
			PlayID:   "test2",
			Name:     "test2",
			TypeName: "comedy",
		},
	}
	performances := []Performance{
		{
			PlayID:   "test",
			Audience: 30,
		},
		{
			PlayID:   "test",
			Audience: 40,
		},
	}
	invoice := Invoice{
		Customer:    "hoge",
		Performance: performances,
	}

	play_service := NewPlayService(plays)

	calc_service := NewTotalAmountCalcService(play_service)

	_, err := calc_service.CalcTotal(&invoice)

	if err != nil {
		t.Errorf("CalcTotal() return error")
	}
}

func Test_CalcTotal_Failed_When_PlayID_Is_Invalid(t *testing.T) {
	plays := []Play{
		{
			PlayID:   "test",
			Name:     "test",
			TypeName: "tragedy",
		},
		{
			PlayID:   "test2",
			Name:     "test2",
			TypeName: "comedy",
		},
	}
	performances := []Performance{
		{
			PlayID:   "hoge",
			Audience: 30,
		},
		{
			PlayID:   "test",
			Audience: 40,
		},
	}
	invoice := Invoice{
		Customer:    "hoge",
		Performance: performances,
	}

	play_service := NewPlayService(plays)

	calc_service := NewTotalAmountCalcService(play_service)

	_, err := calc_service.CalcTotal(&invoice)

	if err == nil {
		t.Errorf("CalcTotal() expected return error. but nil")
	}
}
