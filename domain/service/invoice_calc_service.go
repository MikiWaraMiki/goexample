package domain_service

import (
	"errors"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

type InvoiceCalcService struct {
	play_service *PlayService
}

func (this InvoiceCalcService) GenerateInvoiceReport(invoice *Invoice) (*InvoiceReport, error) {
	totalAmount := NewTotalAmount()
	totalCredit := NewTotalCredit()

	for _, performance := range invoice.Performance {
		play, err := this.play_service.FetchByPlayId(performance.PlayID)

		if err != nil {
			return nil, errors.New("Invalid Play")
		}

		amount := NewAmount(play, &performance)
		totalAmount.AddAmount(amount)

		credit := NewCredit(&performance, play)
		totalCredit.AddCredit(credit)
	}

	invoice_report := NewInvoiceReport(invoice.Customer, totalAmount, totalCredit)

	return invoice_report, nil
}
