package presentation_plain_text

import (
	"fmt"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

type InvoicePlainText struct {
	report *InvoiceReport
}

func NewInvoicePlainText(report *InvoiceReport) *InvoicePlainText {
	return &InvoicePlainText{
		report: report,
	}
}

func (this InvoicePlainText) render() string {
	result := fmt.Sprintf("Statement for %v\n", this.report.CustomerName())

	for _, amount := range this.report.AllAmount() {
		result += fmt.Sprintf("%v: %v (%v seats)\n", amount.PlayName(), amount.FormattedPrice(), amount.Seat())
	}

	result += fmt.Sprintf("Amount owed is %v\n", this.report.TotalAmount())
	result += fmt.Sprintf("You earned %v credits\n", this.report.TotalCredit())

	return result
}
