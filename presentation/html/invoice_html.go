package presentation_html

import (
	"fmt"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

type InvoiceHtml struct {
	report *InvoiceReport
}

func NewInvoiceHtml(report *InvoiceReport) *InvoiceHtml {
	return &InvoiceHtml{
		report: report,
	}
}

func (this InvoiceHtml) Render() string {
	result := fmt.Sprintf(this.h1Statement(), this.report.CustomerName())

	result += "<table>\n"
	result += fmt.Sprintf(this.tableStatement(), "play", "seats", "cost")
	for _, amount := range this.report.AllAmount() {
		result += fmt.Sprintf(this.tableStatement(), amount.PlayName(), amount.Seat(), amount.FormattedPrice())
	}
	result += "</table>\n"
	result += fmt.Sprintf("<p>Amount owned is <em>%v</em></p>\n", this.report.TotalAmount())
	result += fmt.Sprintf("<p>You earned <em>%v</em> credits</p>\n", this.report.TotalCredit())

	return result
}

func (this InvoiceHtml) h1Statement() string {
	return "<h1>Statement for %v"
}

func (this InvoiceHtml) tableStatement() string {
	return "<tr><td>%v</td><td>%v</td><td>%v</td></tr>\n"
}
