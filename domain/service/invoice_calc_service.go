package domain_service

import (
	"errors"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

type InvoiceCalcService struct {
	play_service *PlayService
}

func NewInvoiceCalcService(play_service *PlayService) *InvoiceCalcService {
	return &InvoiceCalcService{
		play_service: play_service,
	}
}

func (this InvoiceCalcService) GenerateInvoiceReport(invoice *Invoice, currency *Currency) (*InvoiceReport, error) {
	totalAmount := NewTotalAmount()
	totalCredit := NewTotalCredit()

	for _, performance := range invoice.Performance {
		play, err := this.play_service.FetchByPlayId(performance.PlayID)

		if err != nil {
			return nil, errors.New("Invalid Play")
		}

		amount := NewAmount(play, &performance, currency)
		totalAmount.AddAmount(amount)

		credit := NewCredit(&performance, play)
		totalCredit.AddCredit(credit)
	}

	invoice_report := NewInvoiceReport(invoice.Customer, totalAmount, totalCredit, currency)

	return invoice_report, nil
}
