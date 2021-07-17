package domain_service

import (
	"errors"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

type TotalCreditCalcService struct {
	play_service *PlayService
}

func NewTotalCreditCalcService(play_service *PlayService) *TotalCreditCalcService {
	return &TotalCreditCalcService{
		play_service: play_service,
	}
}

func (this TotalCreditCalcService) CalcTotal(invoice *Invoice) (*TotalCredit, error) {
	totalCredit := NewTotalCredit()

	for _, performance := range invoice.Performance {
		play, err := this.play_service.FetchByPlayId(performance.PlayID)

		if err != nil {
			return nil, errors.New("Invalid Play")
		}

		credit := NewCredit(&performance, play)

		totalCredit.AddCredit(credit)
	}

	return totalCredit, nil
}
