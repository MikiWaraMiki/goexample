package domain_service

import (
	"errors"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

type TotalAmountCalcService struct {
	play_service *PlayService
}

func NewTotalAmountCalcService(play_service *PlayService) *TotalAmountCalcService {
	return &TotalAmountCalcService{
		play_service: play_service,
	}
}

func (this TotalAmountCalcService) CalcTotal(invoice *Invoice) (*TotalAmount, error) {
	totalAmount := NewTotalAmount()

	for _, performance := range invoice.Performance {
		play, err := this.play_service.FetchByPlayId(performance.PlayID)

		if err != nil {
			return nil, errors.New("Invalid Play")
		}

		amount := NewAmount(play, &performance)

		totalAmount.AddAmount(amount)
	}

	return totalAmount, nil
}
